package router

import (
	"fmt"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/configs"
	notify_handler "github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/handlers/notify"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
	"log/slog"
)

type AMQPRouter struct {
	appcfg *configs.AppConfig
	nh     *notify_handler.NotifyHandler
	cfg    *configs.AMQPConfig
	conn   *rabbitmq.Conn
	cons   *rabbitmq.Consumer
}

func NewAMQPRouter(
	appcfg *configs.AppConfig,
	nh *notify_handler.NotifyHandler,
	cfg *configs.AMQPConfig,
) *AMQPRouter {

	conn, err := rabbitmq.NewConn(
		cfg.URI,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		slog.Error("fail connect to rabbitmq", err)
		return nil
	}

	return &AMQPRouter{
		appcfg: appcfg,
		nh:     nh,
		cfg:    cfg,
		conn:   conn,
	}
}

func (r *AMQPRouter) Setup(app *fiber.App) {
	ncon, err := rabbitmq.NewConsumer(
		r.conn,
		r.cfg.Queue,
		rabbitmq.WithConsumerOptionsConsumerName(fmt.Sprintf("%s:%s", r.appcfg.Name, uuid.New().String())),
		rabbitmq.WithConsumerOptionsRoutingKey("notifier.send-notification"),
		rabbitmq.WithConsumerOptionsExchangeName(r.cfg.Exchange),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsExchangeKind("fanout"),
	)
	if err != nil {
		slog.Error("fail create consumer", err)
		return
	}

	r.cons = ncon
}

func (r *AMQPRouter) Consume() {
	if err := r.cons.Run(r.nh.Handle); err != nil {
		slog.Error("fail run consumer", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}
}

func (r *AMQPRouter) Close() {
	defer r.conn.Close()
	defer r.cons.Close()
}
