package notify_handler

import (
	"context"
	"encoding/json"
	"github.com/WildEgor/e-shop-fiber-microservice-boilerplate/internal/dtos"
	"github.com/stretchr/testify/assert"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"os"
	"testing"
)

type AMQP struct {
	Conn      *rabbitmq.Conn
	Publisher *rabbitmq.Publisher
}

var amqp *AMQP

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		os.Exit(1)
	}

	exitCode := m.Run()

	if err := tearDown(); err != nil {
		os.Exit(1)
	}

	os.Exit(exitCode)
}

func setup() error {

	amqp = &AMQP{}

	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("notifications"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	amqp.Conn = conn
	amqp.Publisher = publisher

	return nil
}

func tearDown() error {
	amqp.Publisher.Close()
	amqp.Conn.Close()
	return nil
}

func TestPublish(t *testing.T) {
	data := &dtos.NotifierPayloadDto{
		Type: "email",
		EmailSetting: dtos.NotifierEmailSettings{
			Email:   "kartashov_egor96@mail.ru",
			Subject: "hi",
			Text:    "<h1>test</h1>",
		},
	}

	bytes, _ := json.Marshal(data)

	err := amqp.Publisher.PublishWithContext(
		context.Background(),
		bytes,
		[]string{"notifier.send-notification"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("notifications"),
	)

	assert.Equal(t, err, nil)
}

func TestPublishLinkTemplate(t *testing.T) {
	data := &dtos.NotifierPayloadDto{
		Type: "email",
		EmailSetting: dtos.NotifierEmailSettings{
			Email:    "kartashov_egor96@mail.ru",
			Subject:  "Confirm code",
			Template: "email_confirm",
			Data: map[string]interface{}{
				"Link": "http://google.com",
			},
		},
	}

	bytes, _ := json.Marshal(data)

	err := amqp.Publisher.PublishWithContext(
		context.Background(),
		bytes,
		[]string{"notifier.send-notification"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("notifications"),
	)

	assert.Equal(t, err, nil)
}

func TestPublishCodeTemplate(t *testing.T) {
	data := &dtos.NotifierPayloadDto{
		Type: "email",
		EmailSetting: dtos.NotifierEmailSettings{
			Email:    "kartashov_egor96@mail.ru",
			Subject:  "Confirm code",
			Template: "email_confirm",
			Data: map[string]interface{}{
				"Code": "1234",
			},
		},
	}

	bytes, _ := json.Marshal(data)

	err := amqp.Publisher.PublishWithContext(
		context.Background(),
		bytes,
		[]string{"notifier.send-notification"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("notifications"),
	)

	assert.Equal(t, err, nil)
}
