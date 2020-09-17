package dao

import (
	"context"
	"errors"
	"github.com/rafaelthomazi/qa/qa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

const (
	// QuestionsCollection is the name of default collection used to store Questions
	QuestionsCollection = "questions"
)

//go:generate mockgen -source=questions.go -destination=questions_mock.go -package=dao

// Questions represents the data access interface to Question object
type Questions interface {
	GetQuestion(id string) (models.Question, error)
	GetQuestions() ([]models.Question, error)
	CreateQuestion(q models.Question) (models.Question, error)
	UpdateQuestion(q models.Question) (models.Question, error)
	DeleteQuestion(id string) error
}

// QuestionsDAO implements the data access interface to Question object
type QuestionsDAO struct {
	client     *mongo.Client
	collection *mongo.Collection
	logger     *zap.Logger
}

// NewQuestionsDAO connects to the MongoDB and instantiates a new QuestionDAO object
func NewQuestionsDAO(cfg Config) (*QuestionsDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	c := client.Database(cfg.Database).Collection(QuestionsCollection)

	return &QuestionsDAO{
		client:     client,
		collection: c,
		logger:     cfg.Logger,
	}, nil
}

// GetQuestion returns a single Question
func (d *QuestionsDAO) GetQuestion(id string) (models.Question, error) {
	logger := d.logger.Named("QuestionsDAO.GetQuestion").With(zap.String("id", id))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	logger.Info("Getting question")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Invalid id", zap.Error(err))
		return models.Question{}, err
	}

	var question models.Question

	filter := bson.M{"_id": oid}

	err = d.collection.FindOne(ctx, filter).Decode(&question)
	if err != nil {
		logger.Error("Error getting question from DB", zap.Error(err))
		return models.Question{}, err
	}

	return question, nil
}

// GetQuestions returns all Questions in the database
func (d *QuestionsDAO) GetQuestions() ([]models.Question, error) {
	logger := d.logger.Named("QuestionsDAO.GetQuestions")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	logger.Info("Getting questions")

	cur, err := d.collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Error("Error getting questions from DB", zap.Error(err))
		return []models.Question{}, err
	}

	defer cur.Close(ctx)

	questions := make([]models.Question, 0)

	for cur.Next(ctx) {
		var q models.Question
		err := cur.Decode(&q)
		if err != nil {
			logger.Error("Error decoding question", zap.Error(err))
			return []models.Question{}, err
		}

		questions = append(questions, q)
	}

	return questions, nil
}

// CreateQuestion inserts a new Question in the database
func (d *QuestionsDAO) CreateQuestion(q models.Question) (models.Question, error) {
	q.ID = primitive.NewObjectID()

	logger := d.logger.Named("QuestionsDAO.CreateQuestion").With(zap.Any("question", q))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	logger.Info("Inserting question")

	_, err := d.collection.InsertOne(ctx, q)
	if err != nil {
		logger.Error("Error creating question in DB", zap.Error(err))
	}

	return q, err
}

// UpdateQuestion updates an existing Question in the database
func (d *QuestionsDAO) UpdateQuestion(q models.Question) (models.Question, error) {
	logger := d.logger.Named("QuestionsDAO.UpdateQuestion")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	filter := bson.M{"_id": q.ID}
	update := bson.M{"$set": q}

	logger.Info("Updating question")

	_, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Error updating question in DB", zap.Error(err))
	}

	return q, nil

}

// DeleteQuestion deletes an existing Question from the database
func (d *QuestionsDAO) DeleteQuestion(id string) error {
	logger := d.logger.Named("QuestionsDAO.DeleteQuestion")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Deleting question")

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Invalid id", zap.Error(err))
		return err
	}

	result, err := d.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		logger.Error("Error deletion question from DB", zap.Error(err))
		return err
	}

	if result.DeletedCount == 0 {
		logger.Error("Question not found in DB")
		return errors.New("question not found")
	}

	return err
}
