package server

import (
	"dominos.com/orders/models"
)

type FindAllOrderItemRequest struct {
	OrderID string
}

type FindAllOrderItemResponse struct {
	OrderItems []models.OrderItem `json:"orderitems"`
	Err        error              `json:"error,omitempty"`
}

func (r FindAllOrderItemResponse) error() error { return r.Err }
