package logs

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints - holds endpoint functions to handle incoming requests.
type Endpoints struct {
	FindAllEndpoint endpoint.Endpoint
	CreateEndpoint  endpoint.Endpoint
}

// MakeEndpoints - returns instance of Endpoints struct
func MakeEndpoints(s TraceLogService) Endpoints {
	return Endpoints{
		FindAllEndpoint: makeFindAllEndpoint(s),
		CreateEndpoint:  makeCreateEndpoint(s),
	}
}

func makeFindAllEndpoint(s TraceLogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		tlogs, e := s.FindAll(ctx)
		return FindAllResponse{Tlogs: tlogs, Err: e}, nil
	}
}

func makeCreateEndpoint(s TraceLogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateRequest)
		e := s.Create(ctx, req.Tlog)
		return CreateResponse{Err: e}, nil
	}
}
