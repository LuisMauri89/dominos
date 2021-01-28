package deliveries

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListarEndPoint    endpoint.Endpoint
	ModificarEndPoint endpoint.Endpoint
	EliminarEndPoint  endpoint.Endpoint
	GetByIdEndpoint   endpoint.Endpoint
}

func MakeEndpoints(s AdvertService) Endpoints {
	return Endpoints{
		ListarEndPoint:    MakeListEndpoint(s),
		ModificarEndPoint: MakeUpdateEndpoint(s),
		EliminarEndPoint:  MakeDeleteEndpoint(s),
		GetByIdEndpoint:   MakeGetByIdEndpoint(s),
	}
}

func MakeListEndpoint(s AdvertService) endpoint.Endpoint { //crear lista de endpoint
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ads, e := s.List(ctx)
		return ListResponse{Ads: ads, Err: e}, nil
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
