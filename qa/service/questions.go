package service

import (
	"context"
	"github.com/rafaelthomazi/qa/qa/models"
)

func (s service) GetQuestion(ctx context.Context, id string) (models.Question, error) {

}

func (s service) GetQuestions(ctx context.Context) ([]models.Question, error) {

}

func (s service) CreateQuestion(ctx context.Context, q models.Question) (models.Question, error) {

}

func (s service) UpdateQuestion(ctx context.Context, q models.Question) (models.Question, error) {

}

func (s service) DeleteQuestion(ctx context.Context, id string) error {

}
