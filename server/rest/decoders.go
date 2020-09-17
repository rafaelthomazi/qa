package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rafaelthomazi/qa/qa/models"
	"net/http"
)

func decodeBlankRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeIDParamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	return id, nil
}

func decodeQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := models.Question{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return models.Question{}, fmt.Errorf("client: %s", err.Error())
	}

	return req, nil
}

func decodeUpdateQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	req := models.Question{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return models.Question{}, fmt.Errorf("client: %s", err.Error())
	}

	if id != req.ID.Hex() {
		return models.Question{}, errors.New("client: URI and Body payload ids do not match")
	}

	return req, nil
}
