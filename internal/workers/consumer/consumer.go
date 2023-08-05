package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WhoIsYgim/swampy_rabbit/internal/config"
	"github.com/WhoIsYgim/swampy_rabbit/internal/pkg/converter"
	"github.com/WhoIsYgim/swampy_rabbit/internal/service"
	swampyrabbit "github.com/WhoIsYgim/swampy_rabbit/pkg/api/message"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitConsumer struct {
	channel  *amqp.Channel
	producer service.Producer
	cfg      *config.RabbitMQConfig
}

func NewConsumer(channel *amqp.Channel, producer service.Producer, cfg *config.RabbitMQConfig) *RabbitConsumer {
	return &RabbitConsumer{
		channel:  channel,
		producer: producer,
		cfg:      cfg,
	}
}

func (c *RabbitConsumer) Process() error {
	return nil
}

func (c *RabbitConsumer) Run() error {

	messages, err := c.channel.Consume(
		c.cfg.QueueCfg.QueueName, // queue name
		"",                       // consumer
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no local
		false,                    // no wait
		nil,                      // arguments
	)
	if err != nil {
		return fmt.Errorf("channel.Consume(): %w", err)
	}

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func(forever chan bool) {
		for message := range messages {
			msg := &swampyrabbit.SwampyRabbit{}
			err := json.Unmarshal(message.Body, msg)
			if err != nil {
				logrus.Warnf("json.Unmarshal(): %v", err)
				return
			}
			send := converter.MessageFromPb(msg.Message)
			err = c.producer.SendMessage(context.Background(), send)
			if err != nil {
				logrus.Warnf("producer.SendMessage(): %v", err)
				return
			}
			logrus.Infof("get a message: %s", msg.Message)
		}
		<-forever
	}(forever)

	return nil
}
