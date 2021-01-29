package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DecodersEncodersOrderItem struct {
	FindAllOrderItemDecoder func(ctx context.Context, r *http.Request) (request interface{}, err error)
	OrderItemEncoder        func(ctx context.Context, w http.ResponseWriter, response interface{}) error
}

func GetDecodersEncodersOrderItem() DecodersEncodersOrderItem {
	return DecodersEncodersOrderItem{
		FindAllOrderItemDecoder: decodeFindAllOrderItemRequest,
		OrderItemEncoder:        encodeOrderItemResponse,
	}
}

func decodeFindAllOrderItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}

	return FindAllOrderItemRequest{OrderID: id}, nil
}

func encodeOrderItemResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
