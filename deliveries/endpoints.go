package deliveries

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FindAllEndpoint endpoint.Endpoint
	CreateEndpoint  endpoint.Endpoint
}

func MakeEndpoints(s DeliveryService) Endpoints {
	return Endpoints{
		FindAllEndpoint: makeFindAllEndpoint(s),
		CreateEndpoint:  makeCreateEndpoint(s),
	}
}

func makeFindAllEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		tlogs, e := s.FindAll(ctx)
		return FindAllResponse{Tlogs: tlogs, Err: e}, nil
	}
}

func makeCreateEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateRequest)
		e := s.Create(ctx, req.Tlog)
		return CreateResponse{Err: e}, nil
	}
}
