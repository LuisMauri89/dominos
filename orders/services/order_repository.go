package services

import (
	"context"
	"errors"

	"dominos.com/orders"
	"dominos.com/orders/models"
	"github.com/jinzhu/gorm"
)

type orderRepository struct {
	conn orders.Connection
}

func NewOrderRepository(conn orders.Connection) OrderRepository {
	return &orderRepository{
		conn: conn,
	}
}

func (r *orderRepository) FindAll(ctx context.Context) ([]models.Order, error) {
	orders := []models.Order{}
	err := r.conn.DB.Find(&orders).Error
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (models.Order, error) {
	order := models.Order{}
	err := r.conn.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.Order{}, errors.New("user not found")
		}
		return models.Order{}, err
	}

	return order, nil
}

func (r *orderRepository) Create(ctx context.Context, order models.Order) error {
	err := r.conn.DB.Create(&order).Error
	if err != nil {
		return err
	}
	r.conn.DB.Save(&order)
	return nil
}

func (r *orderRepository) Delete(ctx context.Context, id string) error {
	err := r.conn.DB.Delete(&models.Order{}, id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
