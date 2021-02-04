package deliveries

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type errorer interface {
	error() error
}

func MakeHTTPHandler(s DeliveryService, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	endpoints := MakeEndpoints(s)
	decodeEncodersOrder := GetDecodersEncodersOrder()
	decodeEncoders := GetDecodersEncoders()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(decodeEncoders.ErrorEncoder),
	}

	router.Methods("GET").Path("/dely/").Handler(httptransport.NewServer(
		endpoints.FindAllEndpoint,
		decodeEncoders.FindAllDecoder,
		decodeEncoders.Encoder,
		options...,
	))

	router.Methods("POST").Path("/dely/").Handler(httptransport.NewServer(
		endpoints.CreateOrderEndpoint,
		decodeEncoders.CreateOrderRequest,
		decodeEncoders.Encoder,
		options...,
	))

	router.Methods("GET").Path("/dely/{status}").Handler(httptransport.NewServer(
		endpoints.GetByStatusEndpoint,
		decodeEncoders.GetByStatusDecoder,
		decodeEncoders.Encoder,
		options...,
	))
	return router
}
