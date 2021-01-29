package repositories

import (
	"context"

	"dominos.com/orders"
	"dominos.com/orders/models"
)

type orderItemRepository struct {
	conn orders.Connection
}

func NewOrderItemRepository(conn orders.Connection) OrderItemRepository {
	return &orderItemRepository{
		conn: conn,
	}
}

func (r *orderItemRepository) FindAll(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	orderitems := []models.OrderItem{}
	err := r.conn.DB.Where("order_id = ?", orderID).Find(&orderitems).Error
	if err != nil {
		return []models.OrderItem{}, err
	}
	return orderitems, nil
}
