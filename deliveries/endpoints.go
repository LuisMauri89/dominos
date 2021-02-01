package deliveries

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FindAllEndpoint     endpoint.Endpoint
	CreateOrderEndpoint endpoint.Endpoint
	GetByStatusEndpoint endpoint.Endpoint
}

func MakeEndpoints(s DeliveryService) Endpoints {
	return Endpoints{
		FindAllEndpoint:     makeFindAllEndpoint(s),
		CreateOrderEndpoint: makeCreateOrderEndpoint(s),
		GetByStatusEndpoint: makeGetByStatusEndpoint(s),
	}
}

func makeFindAllEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		deliveries, e := s.FindAll(ctx)
		return FindAllResponse{TDely: deliveries, Err: e}, nil
	}
}
func makeCreateOrderEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateOrderRequest)
		e := s.Create(ctx, req.order)
		return CreateOrderResponse{Err: e}, nil
	}
}
func makeGetByStatusEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetBySatusRequest)
		deliveries, e := os.GetByID(ctx, req.status)
		return GetByStatusResponse{Order: order, Err: e}, nil
	}
}
