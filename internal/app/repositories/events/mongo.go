package events

import (
	"context"
	"errors"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	database   database.Driver
	collection *mongo.Collection
}

func MakeMongoRepository(database database.Driver) *mongoRepository {
	db := database.GetDatabase().(*mongo.Database)
	collection := db.Collection("events")
	return &mongoRepository{database, collection}
}

func (r *mongoRepository) FindByID(eventID string) (*models.Event, error) {
	if eventID == "" {
		return nil, errors.New("EventID is empty string")
	}

	eventObjectID, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return nil, errors.New("EventID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": eventObjectID})

	event := &models.Event{}

	err = result.Decode(event)

	return event, err
}

func (r *mongoRepository) UpdateByID(eventID string, update interface{}) error {
	if eventID == "" {
		return errors.New("EventID is empty string")
	}

	eventObjectID, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return errors.New("EventID is invalid ObjectID")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": eventObjectID}, &bson.M{"$set": update})

	return err
}

func (r *mongoRepository) DeleteByID(eventID string) error {
	if eventID == "" {
		return errors.New("EventID is empty string")
	}

	eventObjectID, err := primitive.ObjectIDFromHex(eventID)

	if err != nil {
		return errors.New("EventID is invalid ObjectID")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": eventObjectID})

	return err
}

func (r *mongoRepository) Find(filter interface{}) ([]models.Event, error) {
	result, err := r.collection.Find(context.Background(), filter)
	defer result.Close(context.Background())

	events := []models.Event{}

	for result.Next(context.Background()) {
		event := models.Event{}

		result.Decode(&event)

		if result.Err() != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *mongoRepository) FindOne(filter interface{}) (*models.Event, error) {
	result := r.collection.FindOne(context.Background(), filter)

	event := &models.Event{}

	err := result.Decode(event)

	return event, err
}

func (r *mongoRepository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *mongoRepository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *mongoRepository) InsertOne(document models.Event) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *mongoRepository) InsertMany(documents []models.Event) ([]string, error) {
	var i []interface{}

	for _, event := range documents {
		i = append(i, event)
	}

	result, err := r.collection.InsertMany(context.Background(), i)

	insertedObjectIDs := result.InsertedIDs
	insertedIDs := []string{}

	for _, insertedID := range insertedObjectIDs {
		insertedIDs = append(insertedIDs, insertedID.(primitive.ObjectID).Hex())
	}

	return insertedIDs, err
}
