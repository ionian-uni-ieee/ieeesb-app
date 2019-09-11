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
	collection := database.GetCollection("tickets").(*mongod.Collection)
	return &Repository{database, collection}
}

func (r *Repository) FindByID(ticketID string) (*models.Ticket, error) {
	if ticketID == "" {
		return nil, errors.New("TicketID is empty string")
	}

	ticketObjectID, err := primitive.ObjectIDFromHex(ticketID)

	if err != nil {
		return nil, errors.New("TicketID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": ticketObjectID})

	ticket := &models.Ticket{}

	err = result.Decode(ticket)

	return ticket, err
}

func (r *Repository) UpdateByID(ticketID string, update interface{}) error {
	if ticketID == "" {
		return errors.New("TicketID is empty string")
	}

	ticketObjectID, err := primitive.ObjectIDFromHex(ticketID)

	if err != nil {
		return errors.New("TicketID is invalid ObjectID")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": ticketObjectID}, &bson.M{"$set": update})

	return err
}

func (r *Repository) DeleteByID(ticketID string) error {
	if ticketID == "" {
		return errors.New("TicketID is empty string")
	}

	ticketObjectID, err := primitive.ObjectIDFromHex(ticketID)

	if err != nil {
		return errors.New("TicketID is invalid ObjectID")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": ticketObjectID})

	return err
}

func (r *Repository) Find(filter interface{}, skip int64, limit int64) ([]models.Ticket, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.Ticket{}),
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

	tickets := []models.Ticket{}

	for result.Next(context.Background()) {
		ticket := models.Ticket{}

		result.Decode(&ticket)

		if result.Err() != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (r *Repository) FindOne(filter interface{}) (*models.Ticket, error) {
	result := r.collection.FindOne(context.Background(), filter)

	ticket := &models.Ticket{}

	err := result.Decode(ticket)

	return ticket, err
}

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *Repository) InsertOne(document models.Ticket) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *Repository) InsertMany(documents []models.Ticket) ([]string, error) {
	var i []interface{}

	for _, ticket := range documents {
		i = append(i, ticket)
	}

	result, err := r.collection.InsertMany(context.Background(), i)

	insertedObjectIDs := result.InsertedIDs
	insertedIDs := []string{}

	for _, insertedID := range insertedObjectIDs {
		insertedIDs = append(insertedIDs, insertedID.(primitive.ObjectID).Hex())
	}

	return insertedIDs, err
}
