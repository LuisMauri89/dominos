package server

import (
	"dominos.com/orders/models"
)

type FindAllOrderRequest struct {
}

type GetByIDOrderRequest struct {
	ID string `json:"id"`
}

type CreateOrderRequest struct {
	Order models.Order `json:"order"`
}

type DeleteOrderRequest struct {
	ID string `json:"id"`
}

type FindAllOrderResponse struct {
	Orders []models.Order `json:"orders"`
	Err    error          `json:"error,omitempty"`
}

func (r FindAllOrderResponse) error() error { return r.Err }

type GetByIDOrderResponse struct {
	Order models.Order `json:"order"`
	Err   error        `json:"error,omitempty"`
}

func (r GetByIDOrderResponse) error() error { return r.Err }

type CreateOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateOrderResponse) error() error { return r.Err }

type DeleteOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteOrderResponse) error() error { return r.Err }
