package server

import (
	"context"

	"dominos.com/orders/services"
	"github.com/go-kit/kit/endpoint"
)

type OrderItemEndpoints struct {
	FindAllOrderItemEndpoint endpoint.Endpoint
}

func MakeOrderItemEndpoints(ois services.OrderItemService) OrderItemEndpoints {
	return OrderItemEndpoints{
		FindAllOrderItemEndpoint: makeFindAllOrderItemEndpoint(ois),
	}
}

func makeFindAllOrderItemEndpoint(ois services.OrderItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(FindAllOrderItemRequest)
		orderitems, e := ois.FindAll(ctx, req.OrderID)
		return FindAllOrderItemResponse{OrderItems: orderitems, Err: e}, nil
	}
}
