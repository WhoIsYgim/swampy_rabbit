package rabbit

import (
	"fmt"
	"github.com/WhoIsYgim/swampy_rabbit/internal/config"
	"github.com/WhoIsYgim/swampy_rabbit/internal/service"
	"github.com/WhoIsYgim/swampy_rabbit/internal/workers/consumer"
	"github.com/WhoIsYgim/swampy_rabbit/internal/workers/producer"
	"github.com/sirupsen/logrus"
)

func SetupWorkers(cfg *config.RabbitMQConfig) (service.Producer, service.Consumer, error) {
	channel, err := ConnectRabbit(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("rabbit.ConnectRabbit(): %w", err)
	}
	channel.QueueDeclare(
		cfg.QueueCfg.QueueName,
		true,
		true,
		false,
		true,
		nil,
	)
	p := producer.NewRabbitProducer(channel, cfg)
	c := consumer.NewConsumer(channel, p, cfg)
	logrus.Info("workers set up")
	return p, c, nil
}
