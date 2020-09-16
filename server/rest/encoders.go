package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	responseHeaderJSON = "application/json; charset=utf-8"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", responseHeaderJSON)

	if f, ok := response.(Failer); ok && f.Failure() != nil {
		encodeError(ctx, f.Failure(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	statusCode, msg := decodeError(err.Error())
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}

func decodeError(err string) (int, string) {
	switch {
	case strings.HasPrefix(err, "client"):
		return http.StatusBadRequest, strings.TrimPrefix(err, "client: ")
	default:
		return http.StatusInternalServerError, err
	}
}
