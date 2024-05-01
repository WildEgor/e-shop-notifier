package pkg

import (
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/adapters"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/configs"
	eh "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/errors"
	nfm "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/middlewares/not_found"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/router"
	slogger "github.com/WildEgor/e-shop-gopack/pkg/libs/logger/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/google/wire"
	"log/slog"
	"sync"
)

var AppSet = wire.NewSet(
	NewApp,
	router.RouterSet,
	adapters.AdaptersSet,
	configs.ConfigsSet,
)

// Server represents the main server configuration.
type Server struct {
	App        *fiber.App
	AMQPRouter *router.AMQPRouter // TODO: ???
	AppConfig  *configs.AppConfig
}

func (srv *Server) Run() {
	slog.Info(
		"server is listening on PORT",
		slog.String("port", srv.AppConfig.Port))

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		srv.AMQPRouter.Consume()
	}()

	if err := srv.App.Listen(":" + srv.AppConfig.Port); err != nil {
		slog.Error(
			"unable to start server",
			slog.String("error", err.Error()),
		)
		return
	}

	wg.Wait()
}

func (srv *Server) Shutdown() {
	slog.Info(
		"shutdown service",
	)

	srv.AMQPRouter.Close()

	if err := srv.App.Shutdown(); err != nil {
		slog.Error(
			"unable to shutdown server",
			slog.String("error", err.Error()),
		)
	}
}

func NewApp(
	ac *configs.AppConfig,
	eh *eh.ErrorsHandler,
	prr *router.PrivateRouter,
	pbr *router.PublicRouter,
	sr *router.SwaggerRouter,
	ar *router.AMQPRouter,
) *Server {
	logger := slogger.NewLogger(
		slogger.WithAppName(ac.Name),
	)
	if ac.IsProduction() {
		logger = slogger.NewLogger(
			slogger.WithAppName(ac.Name),
			slogger.WithLevel("info"),
			slogger.WithFormat("json"),
		)
	}
	slog.SetDefault(logger)

	app := fiber.New(fiber.Config{
		ErrorHandler: eh.Handle,
		Views:        html.New("./views", ".html"),
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(recover.New())

	prr.Setup(app)
	pbr.Setup(app)
	sr.Setup(app)
	ar.Setup(app)

	// 404 handler
	app.Use(nfm.NewNotFound())

	return &Server{
		App:        app,
		AMQPRouter: ar,
		AppConfig:  ac,
	}
}
