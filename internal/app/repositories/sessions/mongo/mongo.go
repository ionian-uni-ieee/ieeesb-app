package mongo

import (
	"context"
	"errors"
	"reflect"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongod "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	database   database.Driver
	collection *mongod.Collection
}

func MakeRepository(database database.Driver) *Repository {
	collection := database.GetCollection("sessions").(*mongod.Collection)
	return &Repository{database, collection}
}

func (r *Repository) FindByID(sessionID string) (*models.Session, error) {
	if sessionID == "" {
		return nil, errors.New("SessionID is empty string")
	}

	sessionObjectID, err := primitive.ObjectIDFromHex(sessionID)

	if err != nil {
		return nil, errors.New("SessionID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": sessionObjectID})

	session := &models.Session{}

	err = result.Decode(session)

	return session, err
}

func (r *Repository) UpdateByID(sessionID string, update interface{}) error {
	if sessionID == "" {
		return errors.New("SessionID is empty string")
	}

	sessionObjectID, err := primitive.ObjectIDFromHex(sessionID)

	if err != nil {
		return errors.New("SessionID is invalid ObjectID")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": sessionObjectID}, &bson.M{"$set": update})

	return err
}

func (r *Repository) DeleteByID(sessionID string) error {
	if sessionID == "" {
		return errors.New("SessionID is empty string")
	}

	sessionObjectID, err := primitive.ObjectIDFromHex(sessionID)

	if err != nil {
		return errors.New("SessionID is invalid ObjectID")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": sessionObjectID})

	return err
}

func (r *Repository) Find(filter interface{}, skip int64, limit int64) ([]models.Session, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.Session{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	if skip < 0 {
		skip = 0
	}
	if limit <= 0 {
		skip = 12
	}

	result, err := r.collection.Find(
		context.Background(),
		filterBSON,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	defer result.Close(context.Background())

	sessions := []models.Session{}

	for result.Next(context.Background()) {
		session := models.Session{}

		result.Decode(&session)

		if result.Err() != nil {
			return nil, err
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (r *Repository) FindOne(filter interface{}) (*models.Session, error) {
	result := r.collection.FindOne(context.Background(), filter)

	session := &models.Session{}

	err := result.Decode(session)

	return session, err
}

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *Repository) InsertOne(document models.Session) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *Repository) InsertMany(documents []models.Session) ([]string, error) {
	var i []interface{}

	for _, session := range documents {
		i = append(i, session)
	}

	result, err := r.collection.InsertMany(context.Background(), i)

	insertedObjectIDs := result.InsertedIDs
	insertedIDs := []string{}

	for _, insertedID := range insertedObjectIDs {
		insertedIDs = append(insertedIDs, insertedID.(primitive.ObjectID).Hex())
	}

	return insertedIDs, err
}
