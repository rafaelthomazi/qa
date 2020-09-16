package dao

import (
	"context"
	"github.com/rafaelthomazi/qa/qa/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

const (
	QuestionsCollection = "questions"
)

// Questions
type Questions interface {
	GetQuestion(id string) (models.Question, error)
}

type QuestionsDAO struct {
	client     *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewQuestionsDAO(dbURI string, logger *zap.Logger) (*QuestionsDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	c := client.Database("qa").Collection(QuestionsCollection)

	return &QuestionsDAO{
		client:     client,
		collection: c,
		logger:     logger,
	}, nil
}

func (q *QuestionsDAO) GetQuestion(id string) (models.Question, error) {
	logger := q.logger.Named("GetQuestion").With(zap.String("id", id))
}
