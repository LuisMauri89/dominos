package server

import (
	"context"
	"encoding/json"
	"net/http"
)

// DecodersEncoders - holds the necesary functions to decode encode requests and responses.
type DecodersEncoders struct {
	FindAllDecoder func(ctx context.Context, r *http.Request) (request interface{}, err error)
	CreateDecoder  func(ctx context.Context, r *http.Request) (request interface{}, err error)
	UpdateDecoder  func(ctx context.Context, r *http.Request) (request interface{}, err error)
	DeleteDecoder  func(ctx context.Context, r *http.Request) (request interface{}, err error)
	Encoder        func(ctx context.Context, w http.ResponseWriter, response interface{}) error
	ErrorEncoder   func(ctx context.Context, err error, w http.ResponseWriter)
}

// GetDecodersEncoders - returns instance of DecodersEncoders struct
func GetDecodersEncoders() DecodersEncoders {
	return DecodersEncoders{
		FindAllDecoder: decodeFindAllRequest,
		CreateDecoder:  decodeCreateRequest,
		UpdateDecoder:  decodeUpdateRequest,
		DeleteDecoder:  decodeDeleteRequest,
		Encoder:        encodeResponse,
	}
}

func decodeFindAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return FindAllRequest{}, nil
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.TDely); e != nil {
		return nil, e
	}
	return req, nil
}
func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.TDely); e != nil {
		return nil, e
	}
	return req, nil
}
func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.TDely); e != nil {
		return nil, e
	}
	return req, nil
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
		panic("nil error - can not encode nil error.")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
