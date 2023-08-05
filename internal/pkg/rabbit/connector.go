package rabbit

import (
	"fmt"
	"github.com/WhoIsYgim/swampy_rabbit/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func ConnectRabbit(cfg *config.RabbitMQConfig) (*amqp.Channel, error) {
	connectRabbitMQ, err := amqp.Dial(cfg.URl)
	if err != nil {
		return nil, fmt.Errorf("rabbitMQ connect: %w", err)
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil, fmt.Errorf("rabbitMQ connect: %w", err)
	}
	logrus.Info("connected to RabbitMQ")
	return channelRabbitMQ, nil
}
