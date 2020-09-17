package service

import (
	"context"
	"github.com/rafaelthomazi/qa/qa/models"
	"go.uber.org/zap"
)

func (s service) GetQuestion(_ context.Context, id string) (models.Question, error) {
	logger := s.logger.Named("service.GetQuestion").With(zap.String("id", id))

	logger.Info("Getting question")

	question, err := s.questionsDAO.GetQuestion(id)
	if err != nil {
		logger.Error("Error getting question")
		return models.Question{}, err
	}

	return question, nil
}

func (s service) GetQuestions(_ context.Context) ([]models.Question, error) {
	logger := s.logger.Named("service.GetQuestions")

	logger.Info("Getting questions")

	questions, err := s.questionsDAO.GetQuestions()
	if err != nil {
		logger.Error("Error getting questions")
		return []models.Question{}, err
	}

	return questions, nil
}

func (s service) CreateQuestion(_ context.Context, q models.Question) (models.Question, error) {
	logger := s.logger.Named("service.CreateQuestion").With(zap.Any("question", q))

	logger.Info("Creating question")

	question, err := s.questionsDAO.CreateQuestion(q)
	if err != nil {
		logger.Error("Error creating question")
		return models.Question{}, err
	}

	return question, nil
}

func (s service) UpdateQuestion(_ context.Context, q models.Question) (models.Question, error) {
	logger := s.logger.Named("service.UpdateQuestion").With(zap.Any("question", q))

	logger.Info("Updating question")

	question, err := s.questionsDAO.UpdateQuestion(q)
	if err != nil {
		logger.Error("Error updating question")
		return models.Question{}, err
	}

	return question, nil
}

func (s service) DeleteQuestion(_ context.Context, id string) error {
	logger := s.logger.Named("service.DeleteQuestion").With(zap.String("id", id))

	logger.Info("Deleting question")

	err := s.questionsDAO.DeleteQuestion(id)
	if err != nil {
		logger.Error("Error deleting question")
	}

	return err
}
