package deliveries

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrMissingRequiredArguments = errors.New("missing required path variables")
)

func MakeHTTPHandler(s AdvertService, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	endpoints := MakeEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	router.Methods("GET").Path("/send/").Handler(httptransport.NewServer(
		endpoints.ListEndpoint,
		decodeListRequest,
		encodeResponse,
		options...,
	))
	router.Methods("PATCH").Path("/send/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateEndpoint,
		decodeUpdateRequest,
		encodeResponse,
		options...,
	))
	router.Methods("DELETE").Path("/send/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteEndpoint,
		decodeDeleteRequest,
		encodeResponse,
		options...,
	))
	router.Methods("GET").Path("/send/{id}").Handler(httptransport.NewServer(
		endpoints.GetByIdEndpoint,
		decodeGetByIdRequest,
		encodeResponse,
		options...,
	))
	router.Methods("GET").Path("/send/{status}").Handler(httptransport.NewServer(
		endpoints.GetByStatusEndpoint,
		decodeGetByStatusRequest,
		encodeResponse,
		options...,
	))
	return router
}

func decodeListRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return ListRequest{}, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	var advert Ad
	if err := json.NewDecoder(r.Body).Decode(&advert); err != nil {
		return nil, err
	}
	return UpdateRequest{
		ID:     id,
		Advert: advert,
	}, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return DeleteRequest{ID: id}, nil
}

func decodeGetByIdRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return GetByIdRequest{ID: id}, nil
}
func decodeGetByStatusRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["status"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return GetByStatusRequest{Status: status}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
