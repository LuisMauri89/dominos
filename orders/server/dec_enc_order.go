package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DecodersEncodersOrder struct {
	FindAllOrderDecoder func(ctx context.Context, r *http.Request) (request interface{}, err error)
	GetByIDOrderDecoder func(ctx context.Context, r *http.Request) (request interface{}, err error)
	CreateOrderDecoder  func(ctx context.Context, r *http.Request) (request interface{}, err error)
	DeleteOrderDecoder  func(ctx context.Context, r *http.Request) (request interface{}, err error)
	OrderEncoder        func(ctx context.Context, w http.ResponseWriter, response interface{}) error
}

func GetDecodersEncodersOrder() DecodersEncodersOrder {
	return DecodersEncodersOrder{
		FindAllOrderDecoder: decodeFindAllOrderRequest,
		GetByIDOrderDecoder: decodeGetByIDOrderRequest,
		CreateOrderDecoder:  decodeCreateOrderRequest,
		DeleteOrderDecoder:  decodeDeleteOrderRequest,
		OrderEncoder:        encodeOrderResponse,
	}
}

func decodeFindAllOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return FindAllOrderRequest{}, nil
}

func decodeGetByIDOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return GetByIDOrderRequest{ID: id}, nil
}

func decodeCreateOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateOrderRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Order); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return DeleteOrderRequest{ID: id}, nil
}

func encodeOrderResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
