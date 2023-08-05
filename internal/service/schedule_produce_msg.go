package service

import (
	"context"
	"fmt"
	"github.com/WhoIsYgim/swampy_rabbit/internal/pkg/converter"
	swampyservice "github.com/WhoIsYgim/swampy_rabbit/pkg/api/swampy"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SwampyService) ScheduleProduceMsg(ctx context.Context, req *swampyservice.ScheduleProduceMsgRequest) (*emptypb.Empty, error) {
	err := s.producer.SendMessage(ctx, converter.MessageFromPb(req.Message))
	if err != nil {
		return nil, fmt.Errorf("producer.SendMessage(): %w", err)
	}
	return &emptypb.Empty{}, nil
}
