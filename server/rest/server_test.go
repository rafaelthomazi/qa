package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/rafaelthomazi/qa/qa/models"
	"github.com/rafaelthomazi/qa/qa/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_GetQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := models.Question{
		ID:        primitive.NewObjectID(),
		Statement: "test?",
		Answer:    "abc",
	}

	svcMock := service.NewMockService(ctrl)
	svcMock.EXPECT().
		GetQuestion(gomock.Any(), q.ID.Hex()).
		Return(q, nil)

	server := NewServer(svcMock, ":8080", zap.NewNop())

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/qa/questions/%s", q.ID.Hex()), nil)
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := QuestionResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Equal(t, q.ID.Hex(), result.Question.ID.Hex())
	assert.Equal(t, q.Statement, result.Question.Statement)
	assert.Equal(t, q.Answer, result.Question.Answer)
}

func TestServer_GetQuestionNoFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	qid := "test"

	svcMock := service.NewMockService(ctrl)
	svcMock.EXPECT().
		GetQuestion(gomock.Any(), "test").
		Return(models.Question{}, errors.New("not found"))

	server := NewServer(svcMock, ":8080", zap.NewNop())

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/qa/questions/%s", qid), nil)
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	result := ErrorResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Equal(t, "not found", result.Error)
}

func TestServer_GetQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	questions := []models.Question{
		{
			ID:        primitive.NewObjectID(),
			Statement: "test?",
			Answer:    "abc",
		},
	}

	svcMock := service.NewMockService(ctrl)
	svcMock.EXPECT().
		GetQuestions(gomock.Any()).
		Return(questions, nil)

	server := NewServer(svcMock, ":8080", zap.NewNop())

	req := httptest.NewRequest("GET", "/api/qa/questions", nil)
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := QuestionsResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Len(t, result.Questions, 1)
	assert.Equal(t, questions[0].ID.Hex(), result.Questions[0].ID.Hex())
}

func TestServer_CreateQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := models.Question{
		Statement: "test?",
	}

	svcMock := service.NewMockService(ctrl)
	svcMock.EXPECT().
		CreateQuestion(gomock.Any(), q).
		Return(q, nil)

	server := NewServer(svcMock, ":8080", zap.NewNop())

	qStr, err := json.Marshal(q)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/qa/questions", strings.NewReader(string(qStr)))
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := QuestionResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Question.ID.Hex())
	assert.Equal(t, q.Statement, result.Question.Statement)
}

func TestServer_UpdateQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := models.Question{
		ID:        primitive.NewObjectID(),
		Statement: "test?",
	}

	svcMock := service.NewMockService(ctrl)
	svcMock.EXPECT().
		UpdateQuestion(gomock.Any(), q).
		Return(q, nil)

	server := NewServer(svcMock, ":8080", zap.NewNop())

	qStr, err := json.Marshal(q)
	assert.NoError(t, err)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/qa/questions/%s", q.ID.Hex()), strings.NewReader(string(qStr)))
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := QuestionResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Equal(t, q.ID.Hex(), result.Question.ID.Hex())
	assert.Equal(t, q.Statement, result.Question.Statement)
}

func TestServer_UpdateQuestionInvalidPayload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	q := models.Question{
		ID:        primitive.NewObjectID(),
		Statement: "test?",
	}

	svcMock := service.NewMockService(ctrl)
	server := NewServer(svcMock, ":8080", zap.NewNop())

	qStr, err := json.Marshal(q)
	assert.NoError(t, err)

	req := httptest.NewRequest("PUT", "/api/qa/questions/1234", strings.NewReader(string(qStr)))
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	result := ErrorResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Equal(t, "URI and Body payload ids do not match", result.Error)
}

func TestServer_DeleteQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	qid := primitive.NewObjectID()

	svcMock := service.NewMockService(ctrl)
	svcMock.EXPECT().
		DeleteQuestion(gomock.Any(), qid.Hex()).
		Return(nil)

	server := NewServer(svcMock, ":8080", zap.NewNop())

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/qa/questions/%s", qid.Hex()), nil)
	w := httptest.NewRecorder()

	server.Handler.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := StringResponse{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.Equal(t, qid.Hex(), result.Value)
}
