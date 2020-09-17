package service

import (
	"context"
	"github.com/rafaelthomazi/qa/qa/dao"
	"github.com/rafaelthomazi/qa/qa/models"
	"go.uber.org/zap"
)

// Service represents the interface "qa" service will implements
type Service interface {
	GetQuestion(ctx context.Context, id string) (models.Question, error)
	GetQuestions(ctx context.Context) ([]models.Question, error)
	CreateQuestion(ctx context.Context, q models.Question) (models.Question, error)
	UpdateQuestion(ctx context.Context, q models.Question) (models.Question, error)
	DeleteQuestion(ctx context.Context, id string) error
}

type service struct {
	questionsDAO dao.Questions
	logger       *zap.Logger
}

// NewService ...
func NewService(cfg *Config) Service {
	questionsDAO, err := dao.NewQuestionsDAO(cfg.DAO)
	if err != nil {
		cfg.Logger.Error("Error creating Questions DAO", zap.Error(err))
		return nil
	}

	return &service{
		questionsDAO: questionsDAO,
		logger:       cfg.Logger,
	}
}
