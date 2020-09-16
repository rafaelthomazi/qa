package rest

import (
	"context"
	"github.com/rafaelthomazi/qa/qa/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/rafaelthomazi/qa/qa/service"
	"go.uber.org/zap"
)

// Endpoints is the complete list of all service calls that can be wired up to HTTP or GRPC
type Endpoints struct {
	GetQuestionEndpoint    endpoint.Endpoint
	GetQuestionsEndpoint   endpoint.Endpoint
	CreateQuestionEndpoint endpoint.Endpoint
	UpdateQuestionEndpoint endpoint.Endpoint
	DeleteQuestionEndpoint endpoint.Endpoint
}

// New returns all callable service actions available to REST server
func New(svc service.Service, logger *zap.Logger) Endpoints {
	return Endpoints{
		GetQuestionEndpoint:    makeGetQuestionEndpoint(svc, logger),
		GetQuestionsEndpoint:   makeGetQuestionsEndpoint(svc, logger),
		CreateQuestionEndpoint: makeCreateQuestionEndpoint(svc, logger),
		UpdateQuestionEndpoint: makeUpdateQuestionEndpoint(svc, logger),
		DeleteQuestionEndpoint: makeDeleteQuestionEndpoint(svc, logger),
	}
}

func makeGetQuestionEndpoint(svc service.Service, logger *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qID := request.(string)
		logger.Debug("Request GetQuestion", zap.String("question_id", qID))
		result, err := svc.GetQuestion(ctx, qID)
		logger.Debug("Response", zap.Any("question", result), zap.Error(err))
		return QuestionResponse{result, err}, nil
	}
}

func makeGetQuestionsEndpoint(svc service.Service, logger *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger.Debug("Request GetQuestions")
		result, err := svc.GetQuestions(ctx)
		logger.Debug("Response", zap.Any("questions", result), zap.Error(err))
		return QuestionsResponse{result, err}, nil
	}
}

func makeCreateQuestionEndpoint(svc service.Service, logger *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		q := request.(models.Question)
		logger.Debug("Request CreateQuestion", zap.Any("question", q))
		result, err := svc.CreateQuestion(ctx, q)
		logger.Debug("Response", zap.Any("question", result), zap.Error(err))
		return QuestionResponse{result, err}, nil
	}
}

func makeUpdateQuestionEndpoint(svc service.Service, logger *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		q := request.(models.Question)
		logger.Debug("Request UpdateQuestion", zap.Any("question", q))
		result, err := svc.UpdateQuestion(ctx, q)
		logger.Debug("Response", zap.Any("question", result), zap.Error(err))
		return QuestionResponse{result, err}, nil
	}
}

func makeDeleteQuestionEndpoint(svc service.Service, logger *zap.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		qID := request.(string)
		logger.Debug("Request DeleteQuestion", zap.String("question_id", qID))
		err := svc.DeleteQuestion(ctx, qID)
		logger.Debug("Response", zap.Error(err))
		return StringResponse{qID, err}, nil
	}
}
