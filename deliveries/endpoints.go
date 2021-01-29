package deliveries

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FindAllEndpoint endpoint.Endpoint
}

func MakeEndpoints(s DeliveryService) Endpoints {
	return Endpoints{
		FindAllEndpoint: makeFindAllEndpoint(s),
	}
}

func makeFindAllEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		deliveries, e := s.FindAll(ctx)
		return FindAllResponse{TDely: deliveries, Err: e}, nil
	}
}
