package server

import (
	"context"

	"dominos.com/orders/services"
	"github.com/go-kit/kit/endpoint"
)

type OrderEndpoints struct {
	FindAllOrderEndpoint endpoint.Endpoint
	GetByIDOrderEndpoint endpoint.Endpoint
	CreateOrderEndpoint  endpoint.Endpoint
	DeleteOrderEndpoint  endpoint.Endpoint
}

func MakeOrderEndpoints(os services.OrderService) OrderEndpoints {
	return OrderEndpoints{
		FindAllOrderEndpoint: makeFindAllOrderEndpoint(os),
		GetByIDOrderEndpoint: makeGetByIDOrderEndpoint(os),
		CreateOrderEndpoint:  makeCreateOrderEndpoint(os),
		DeleteOrderEndpoint:  makeDeleteOrderEndpoint(os),
	}
}

func makeFindAllOrderEndpoint(os services.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		orders, e := os.FindAll(ctx)
		return FindAllOrderResponse{Orders: orders, Err: e}, nil
	}
}

func makeGetByIDOrderEndpoint(os services.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetByIDOrderRequest)
		order, e := os.GetByID(ctx, req.ID)
		return GetByIDOrderResponse{Order: order, Err: e}, nil
	}
}

func makeCreateOrderEndpoint(os services.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateOrderRequest)
		e := os.Create(ctx, req.Order)
		return CreateOrderResponse{Err: e}, nil
	}
}

func makeDeleteOrderEndpoint(os services.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteOrderRequest)
		e := os.Delete(ctx, req.ID)
		return DeleteOrderResponse{Err: e}, nil
	}
}
