package service

import (
	"context"
	"github.com/WhoIsYgim/swampy_rabbit/internal/domain"
	swampyservice "github.com/WhoIsYgim/swampy_rabbit/pkg/api/swampy"
)

type Producer interface {
	SendMessage(ctx context.Context, message *domain.Message) error
}

type Consumer interface {
	Process() error
	Run() error
}

type SwampyService struct {
	swampyservice.UnimplementedSwampyServiceServer

	producer Producer
	consumer Consumer
}

// NewSwampyService - создает новый сервис
func NewSwampyService(producer Producer, consumer Consumer) *SwampyService {
	return &SwampyService{
		producer: producer,
		consumer: consumer,
	}
}
