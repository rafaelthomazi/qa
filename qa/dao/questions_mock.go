// Code generated by MockGen. DO NOT EDIT.
// Source: questions.go

// Package dao is a generated GoMock package.
package dao

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/rafaelthomazi/qa/qa/models"
	reflect "reflect"
)

// MockQuestions is a mock of Questions interface
type MockQuestions struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionsMockRecorder
}

// MockQuestionsMockRecorder is the mock recorder for MockQuestions
type MockQuestionsMockRecorder struct {
	mock *MockQuestions
}

// NewMockQuestions creates a new mock instance
func NewMockQuestions(ctrl *gomock.Controller) *MockQuestions {
	mock := &MockQuestions{ctrl: ctrl}
	mock.recorder = &MockQuestionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuestions) EXPECT() *MockQuestionsMockRecorder {
	return m.recorder
}

// GetQuestion mocks base method
func (m *MockQuestions) GetQuestion(id string) (models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestion", id)
	ret0, _ := ret[0].(models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestion indicates an expected call of GetQuestion
func (mr *MockQuestionsMockRecorder) GetQuestion(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestion", reflect.TypeOf((*MockQuestions)(nil).GetQuestion), id)
}

// GetQuestions mocks base method
func (m *MockQuestions) GetQuestions() ([]models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestions")
	ret0, _ := ret[0].([]models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestions indicates an expected call of GetQuestions
func (mr *MockQuestionsMockRecorder) GetQuestions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestions", reflect.TypeOf((*MockQuestions)(nil).GetQuestions))
}

// CreateQuestion mocks base method
func (m *MockQuestions) CreateQuestion(q models.Question) (models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuestion", q)
	ret0, _ := ret[0].(models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuestion indicates an expected call of CreateQuestion
func (mr *MockQuestionsMockRecorder) CreateQuestion(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestion", reflect.TypeOf((*MockQuestions)(nil).CreateQuestion), q)
}

// UpdateQuestion mocks base method
func (m *MockQuestions) UpdateQuestion(q models.Question) (models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuestion", q)
	ret0, _ := ret[0].(models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuestion indicates an expected call of UpdateQuestion
func (mr *MockQuestionsMockRecorder) UpdateQuestion(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuestion", reflect.TypeOf((*MockQuestions)(nil).UpdateQuestion), q)
}

// DeleteQuestion mocks base method
func (m *MockQuestions) DeleteQuestion(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuestion", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuestion indicates an expected call of DeleteQuestion
func (mr *MockQuestionsMockRecorder) DeleteQuestion(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestion", reflect.TypeOf((*MockQuestions)(nil).DeleteQuestion), id)
}