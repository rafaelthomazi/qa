package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/rafaelthomazi/qa/qa/dao"
	"github.com/rafaelthomazi/qa/qa/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"testing"
)

func TestService_GetQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	qid := primitive.NewObjectID()
	q := models.Question{
		ID:        qid,
		Statement: "test?",
		Answer:    "abc",
	}

	daoMock := dao.NewMockQuestions(ctrl)
	daoMock.EXPECT().
		GetQuestion(qid.Hex()).
		Return(q, nil)

	svc := service{
		questionsDAO: daoMock,
		logger:       zap.NewNop(),
	}

	question, err := svc.GetQuestion(context.Background(), qid.Hex())
	assert.NoError(t, err)
	assert.Equal(t, qid.Hex(), question.ID.Hex())
	assert.Equal(t, "test?", question.Statement)
	assert.Equal(t, "abc", question.Answer)
}

func TestService_GetQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	questions := []models.Question{
		{
			ID:        primitive.NewObjectID(),
			Statement: "test?",
			Answer:    "abc",
		},
	}

	daoMock := dao.NewMockQuestions(ctrl)
	daoMock.EXPECT().
		GetQuestions().
		Return(questions, nil)

	svc := service{
		questionsDAO: daoMock,
		logger:       zap.NewNop(),
	}

	result, err := svc.GetQuestions(context.Background())
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestService_CreateQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	question := models.Question{
		Statement: "test?",
	}

	daoMock := dao.NewMockQuestions(ctrl)
	daoMock.EXPECT().
		CreateQuestion(question).
		Return(
			models.Question{
				ID:        primitive.NewObjectID(),
				Statement: "test?",
			}, nil)

	svc := service{
		questionsDAO: daoMock,
		logger:       zap.NewNop(),
	}

	result, err := svc.CreateQuestion(context.Background(), question)
	assert.NoError(t, err)
	assert.False(t, result.ID.IsZero())
	assert.Equal(t, "test?", result.Statement)
}

func TestService_UpdateQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	question := models.Question{
		ID:        primitive.NewObjectID(),
		Statement: "test?",
		Answer:    "abc",
	}

	daoMock := dao.NewMockQuestions(ctrl)
	daoMock.EXPECT().
		UpdateQuestion(question).
		Return(question, nil)

	svc := service{
		questionsDAO: daoMock,
		logger:       zap.NewNop(),
	}

	result, err := svc.UpdateQuestion(context.Background(), question)
	assert.NoError(t, err)
	assert.Equal(t, question.ID, result.ID)
	assert.Equal(t, question.Statement, result.Statement)
	assert.Equal(t, question.Answer, result.Answer)
}

func TestService_DeleteQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	qid := primitive.NewObjectID().Hex()

	daoMock := dao.NewMockQuestions(ctrl)
	daoMock.EXPECT().
		DeleteQuestion(qid).
		Return(nil)

	svc := service{
		questionsDAO: daoMock,
		logger:       zap.NewNop(),
	}

	err := svc.DeleteQuestion(context.Background(), qid)
	assert.NoError(t, err)
}
