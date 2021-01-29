package deliveries

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type errorer interface {
	error() error
}

// MakeHTTPHandler -
func MakeHTTPHandler(s TraceLogService, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	endpoints := MakeEndpoints(s)
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
		endpoints.CreateEndpoint,
		decodeEncoders.CreateDecoder,
		decodeEncoders.Encoder,
		options...,
	))
	router.Methods("PUT").Path("/dely/").Handler(httptransport.NewServer(
		endpoints.UpdateEndpoint,
		decodeEncoders.UpdateDecoder,
		decodeEncoders.Encoder,
		options...,
	))
	router.Methods("DELETE").Path("/dely/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteEndpoint,
		decodeEncoders.DeleteDecoder,
		decodeEncoders.Encoder,
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
