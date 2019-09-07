package mongo

import (
	"context"
	"errors"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongod "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database   database.Driver
	collection *mongod.Collection
}

func MakeRepository(database database.Driver) *Repository {
	collection := database.GetCollection("events").(*mongod.Collection)
	return &Repository{database, collection}
}

func (r *Repository) FindByID(eventID string) (*models.Event, error) {
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

func (r *Repository) UpdateByID(eventID string, update interface{}) error {
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

func (r *Repository) DeleteByID(eventID string) error {
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

func (r *Repository) Find(filter interface{}) ([]models.Event, error) {
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

func (r *Repository) FindOne(filter interface{}) (*models.Event, error) {
	result := r.collection.FindOne(context.Background(), filter)

	event := &models.Event{}

	err := result.Decode(event)

	return event, err
}

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *Repository) InsertOne(document models.Event) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *Repository) InsertMany(documents []models.Event) ([]string, error) {
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

func (r *Repository) IsDuplicate(name string) bool {
	sameKeysFilter := &bson.M{
		"name": name,
	}

	eventFound := models.Sponsor{}
	r.collection.FindOne(context.Background(), sameKeysFilter).Decode(&eventFound)

	return !eventFound.ID.IsZero()
}
