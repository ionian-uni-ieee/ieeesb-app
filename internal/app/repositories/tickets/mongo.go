package tickets

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
	collection := db.Collection("tickets")
	return &mongoRepository{database, collection}
}

func (r *mongoRepository) FindByID(ticketID string) (*models.Ticket, error) {
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

func (r *mongoRepository) UpdateByID(ticketID string, update interface{}) error {
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

func (r *mongoRepository) DeleteByID(ticketID string) error {
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

func (r *mongoRepository) Find(filter interface{}) ([]models.Ticket, error) {
	result, err := r.collection.Find(context.Background(), filter)
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

func (r *mongoRepository) FindOne(filter interface{}) (*models.Ticket, error) {
	result := r.collection.FindOne(context.Background(), filter)

	ticket := &models.Ticket{}

	err := result.Decode(ticket)

	return ticket, err
}

func (r *mongoRepository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *mongoRepository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *mongoRepository) InsertOne(document models.Ticket) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *mongoRepository) InsertMany(documents []models.Ticket) ([]string, error) {
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
