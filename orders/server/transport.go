package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"dominos.com/orders/services"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrMissingRequiredArguments = errors.New("missing required argument")
)

type errorer interface {
	error() error
}

// MakeHTTPHandler -
func MakeHTTPHandler(os services.OrderService, ois services.OrderItemService, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	orderEndpoints := MakeOrderEndpoints(os)
	orderItemEndpoints := MakeOrderItemEndpoints(ois)
	decodeEncodersOrder := GetDecodersEncodersOrder()
	decodeEncodersOrderItem := GetDecodersEncodersOrderItem()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(EncodeError),
	}

	router.Methods("GET").Path("/orders/").Handler(httptransport.NewServer(
		orderEndpoints.FindAllOrderEndpoint,
		decodeEncodersOrder.FindAllOrderDecoder,
		decodeEncodersOrder.OrderEncoder,
		options...,
	))
	router.Methods("GET").Path("/orders/{id}").Handler(httptransport.NewServer(
		orderEndpoints.GetByIDOrderEndpoint,
		decodeEncodersOrder.GetByIDOrderDecoder,
		decodeEncodersOrder.OrderEncoder,
		options...,
	))
	router.Methods("POST").Path("/orders/").Handler(httptransport.NewServer(
		orderEndpoints.CreateOrderEndpoint,
		decodeEncodersOrder.CreateOrderDecoder,
		decodeEncodersOrder.OrderEncoder,
		options...,
	))
	router.Methods("DELETE").Path("/orders/{id}").Handler(httptransport.NewServer(
		orderEndpoints.DeleteOrderEndpoint,
		decodeEncodersOrder.DeleteOrderDecoder,
		decodeEncodersOrder.OrderEncoder,
		options...,
	))
	router.Methods("GET").Path("/order/{id}/orderitems/").Handler(httptransport.NewServer(
		orderItemEndpoints.FindAllOrderItemEndpoint,
		decodeEncodersOrderItem.FindAllOrderItemDecoder,
		decodeEncodersOrderItem.OrderItemEncoder,
		options...,
	))
	return router
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("nil error - can not encode nil error.")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
