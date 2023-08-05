package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WhoIsYgim/swampy_rabbit/internal/config"
	"github.com/WhoIsYgim/swampy_rabbit/internal/domain"
	"github.com/streadway/amqp"
)

type RabbitProducer struct {
	channel *amqp.Channel
	cfg     *config.RabbitMQConfig
}

func NewRabbitProducer(channel *amqp.Channel, cfg *config.RabbitMQConfig) *RabbitProducer {
	return &RabbitProducer{
		channel: channel,
		cfg:     cfg,
	}
}

func (p *RabbitProducer) SendMessage(ctx context.Context, message *domain.Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("json.Marshall(): %w", err)
	}
	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	err = p.channel.Publish(
		p.cfg.QueueCfg.Exchange,
		p.cfg.QueueCfg.QueueName,
		false,
		false,
		msg,
	)
	if err != nil {
		return fmt.Errorf("channel.Publish(): %w", err)
	}
	return nil
}
