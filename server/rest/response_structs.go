package rest

import "github.com/rafaelthomazi/qa/qa/models"

// Failer provides a common failed response mechanism
type Failer interface {
	Failure() error
}

// StringResponse ..
type StringResponse struct {
	Value string `json:"value"`
	Error error  `json:"-"`
}

// Failure returns the error if any
func (r StringResponse) Failure() error { return r.Error }

// QuestionResponse ..
type QuestionResponse struct {
	Question models.Question `json:"question"`
	Error    error           `json:"-"`
}

// Failure returns the error if any
func (r QuestionResponse) Failure() error { return r.Error }

// QuestionsResponse ..
type QuestionsResponse struct {
	Questions []models.Question `json:"questions"`
	Error     error             `json:"-"`
}

// Failure returns the error if any
func (r QuestionsResponse) Failure() error { return r.Error }

// ErrorResponse ..
type ErrorResponse struct {
	Error string `json:"error"`
}
