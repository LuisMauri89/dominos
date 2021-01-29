package services

import (
	"context"

	"dominos.com/orders/models"
	"dominos.com/orders/repositories"
	"github.com/go-kit/kit/log"
)

type orderService struct {
	repository repositories.OrderRepository
	logger     log.Logger
}

func NewOrderService(repository repositories.OrderRepository, logger log.Logger) OrderService {
	return &orderService{
		repository: repository,
		logger:     logger,
	}
}

func (s *orderService) FindAll(ctx context.Context) ([]models.Order, error) {
	orders, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return orders, nil
}

func (s *orderService) GetByID(ctx context.Context, id string) (models.Order, error) {
	order, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return order, err
	}
	s.logger.Log("getbyid:", "success")
	return order, nil
}

func (s *orderService) Create(ctx context.Context, order models.Order) error {
	order.Prepare()
	err := order.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	items := []models.OrderItem{}
	for _, orderitem := range order.Items {
		orderitem.Prepare()
		err = orderitem.Validate()
		if err != nil {
			s.logger.Log("error:", err)
			return err
		}
		items = append(items, orderitem)
	}
	order.Items = items

	if err := s.repository.Create(ctx, order); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")
	return nil
}

func (s *orderService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
