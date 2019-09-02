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
	collection := database.GetCollection("sponsors").(*mongod.Collection)
	return &Repository{database, collection}
}

func (r *Repository) FindByID(sponsorID string) (*models.Sponsor, error) {
	if sponsorID == "" {
		return nil, errors.New("SponsorID is empty string")
	}

	sponsorObjectID, err := primitive.ObjectIDFromHex(sponsorID)

	if err != nil {
		return nil, errors.New("SponsorID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": sponsorObjectID})

	sponsor := &models.Sponsor{}

	err = result.Decode(sponsor)

	return sponsor, err
}

func (r *Repository) UpdateByID(sponsorID string, update interface{}) error {
	if sponsorID == "" {
		return errors.New("SponsorID is empty string")
	}

	sponsorObjectID, err := primitive.ObjectIDFromHex(sponsorID)

	if err != nil {
		return errors.New("SponsorID is invalid ObjectID")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": sponsorObjectID}, &bson.M{"$set": update})

	return err
}

func (r *Repository) DeleteByID(sponsorID string) error {
	if sponsorID == "" {
		return errors.New("SponsorID is empty string")
	}

	sponsorObjectID, err := primitive.ObjectIDFromHex(sponsorID)

	if err != nil {
		return errors.New("SponsorID is invalid ObjectID")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": sponsorObjectID})

	return err
}

func (r *Repository) Find(filter interface{}) ([]models.Sponsor, error) {
	result, err := r.collection.Find(context.Background(), filter)
	defer result.Close(context.Background())

	sponsors := []models.Sponsor{}

	for result.Next(context.Background()) {
		sponsor := models.Sponsor{}

		result.Decode(&sponsor)

		if result.Err() != nil {
			return nil, err
		}

		sponsors = append(sponsors, sponsor)
	}

	return sponsors, nil
}

func (r *Repository) FindOne(filter interface{}) (*models.Sponsor, error) {
	result := r.collection.FindOne(context.Background(), filter)

	sponsor := &models.Sponsor{}

	err := result.Decode(sponsor)

	return sponsor, err
}

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *Repository) InsertOne(document models.Sponsor) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *Repository) InsertMany(documents []models.Sponsor) ([]string, error) {
	var i []interface{}

	for _, sponsor := range documents {
		i = append(i, sponsor)
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
	return false
}
