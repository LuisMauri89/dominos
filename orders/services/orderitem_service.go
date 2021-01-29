package services

import (
	"context"

	"dominos.com/orders/models"
	"github.com/go-kit/kit/log"
)

type orderItemService struct {
	repository OrderItemRepository
	logger     log.Logger
}

func NewOrderItemService(repository OrderItemRepository, logger log.Logger) OrderItemService {
	return &orderItemService{
		repository: repository,
		logger:     logger,
	}
}

func (s *orderItemService) FindAll(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	orderitems, err := s.repository.FindAll(ctx, orderID)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return orderitems, nil
}
