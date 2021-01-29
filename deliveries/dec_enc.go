package deliveries

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

// DecodersEncoders - holds the necesary functions to decode encode requests and responses.
type DecodersEncoders struct {
	FindAllDecoder func(ctx context.Context, r *http.Request) (request interface{}, err error)
	Encoder        func(ctx context.Context, w http.ResponseWriter, response interface{}) error
	ErrorEncoder   func(ctx context.Context, err error, w http.ResponseWriter)
}

// GetDecodersEncoders - returns instance of DecodersEncoders struct
func GetDecodersEncoders() DecodersEncoders {
	return DecodersEncoders{
		FindAllDecoder: decodeFindAllRequest,
		Encoder:        encodeResponse,
		ErrorEncoder:   encodeError,
	}
}

func decodeFindAllRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return FindAllRequest{}, nil
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

func codeFrom(err error) int {
	switch err {
	case ErrNotFound, sql.ErrNoRows:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
