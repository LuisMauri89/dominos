package services

import (
	"context"

	"dominos.com/orders/models"
)

type OrderRepository interface {
	FindAll(ctx context.Context) ([]models.Order, error)
	GetByID(ctx context.Context, id string) (models.Order, error)
	Create(ctx context.Context, order models.Order) error
	Delete(ctx context.Context, id string) error
}

type OrderItemRepository interface {
	FindAll(ctx context.Context, orderID string) ([]models.OrderItem, error)
}
