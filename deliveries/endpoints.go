package DeliveriesV2

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FindAllEndpoint endpoint.Endpoint
	UpdateEndpoint  endpoint.Endpoint
	DeleteEndpoint  endpoint.Endpoint
	GetByIdEndpoint endpoint.Endpoint
}

func MakeEndpoints(s TraceLogService) Endpoints {
	return Endpoints{
		FindAllEndpoint: makeFindAllEndpoint(s),
		UpdateEndpoint:  MakeUpdateEndpointEndpoint(s),
		DeleteEndpoint:  MakeDeleteEndpointEndpoint(s),
		GetByIdEndpoint: MakeGetByIdEndpointEndpoint(s),
	}
}

func makeFindAllEndpoint(s AdvertService) endpoint.Endpoint { //crear lista de endpoint
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		Dely, e := s.List(ctx)
		return ListResponse{Dely: dely, Err: e}, nil
	}
}

func MakeUpdateEndpoint(s AdvertService) endpoint.Endpoint { // modificar endpoint
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateRequest)
		e := s.Update(ctx, req.ID, req.Advert)
		return UpdateResponse{Err: e}, nil
	}
}

func MakeDeleteEndpoint(s AdvertService) endpoint.Endpoint { //eliminar endpoint
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteRequest)
		e := s.Delete(ctx, req.ID)
		return DeleteResponse{Err: e}, nil
	}
}

func MakeGetByIdEndpoint(s AdvertService) endpoint.Endpoint { //busqueda por id endpoint
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetByIdRequest)
		ad, e := s.GetById(ctx, req.ID)
		return GetByIdResponse{Advert: ad, Err: e}, nil
	}
}
