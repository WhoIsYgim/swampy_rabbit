package converter

import (
	"github.com/WhoIsYgim/swampy_rabbit/internal/domain"
	"time"
)

func MessageFromPb(msg string) *domain.Message {
	return &domain.Message{
		Payload:   msg,
		CreatedAt: time.Now(),
	}
}
