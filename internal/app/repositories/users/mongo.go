package users

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
	collection := db.Collection("users")
	return &mongoRepository{database, collection}
}

func (r *mongoRepository) FindByID(userID string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("UserID is empty string")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, errors.New("UserID is invalid ObjectID")
	}

	result := r.collection.FindOne(context.Background(), &bson.M{"_id": userObjectID})

	user := &models.User{}

	err = result.Decode(user)

	return user, err
}

func (r *mongoRepository) UpdateByID(userID string, update interface{}) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return errors.New("UserID is invalid ObjectID")
	}

	user := r.collection.FindOne(context.Background(), &bson.M{"_id": userObjectID})

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("No user with this ID exists")
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": userObjectID}, &bson.M{"$set": update})

	return err
}

func (r *mongoRepository) DeleteByID(userID string) error {
	if userID == "" {
		return errors.New("UserID is empty string")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return errors.New("UserID is invalid ObjectID")
	}

	user := r.collection.FindOne(context.Background(), &bson.M{"_id": userObjectID})

	if user == nil {
		return errors.New("No user with this ID exists")
	}

	_, err = r.collection.DeleteOne(context.Background(), &bson.M{"_id": userObjectID})

	return err
}

func (r *mongoRepository) Find(filter interface{}) ([]models.User, error) {
	result, err := r.collection.Find(context.Background(), filter)
	defer result.Close(context.Background())

	users := []models.User{}

	for result.Next(context.Background()) {
		user := models.User{}

		result.Decode(&user)

		if result.Err() != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *mongoRepository) FindOne(filter interface{}) (*models.User, error) {
	result := r.collection.FindOne(context.Background(), filter)

	user := &models.User{}

	err := result.Decode(user)

	return user, err
}

func (r *mongoRepository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	result, err := r.collection.UpdateMany(context.Background(), filter, update)

	return result.UpsertedID.([]string), err
}

func (r *mongoRepository) DeleteMany(filter interface{}) (int64, error) {
	result, err := r.collection.DeleteMany(context.Background(), filter)

	return result.DeletedCount, err
}

func (r *mongoRepository) InsertOne(document models.User) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *mongoRepository) InsertMany(documents []models.User) ([]string, error) {
	var i []interface{}

	for _, user := range documents {
		i = append(i, user)
	}

	result, err := r.collection.InsertMany(context.Background(), i)

	insertedObjectIDs := result.InsertedIDs
	insertedIDs := []string{}

	for _, insertedID := range insertedObjectIDs {
		insertedIDs = append(insertedIDs, insertedID.(primitive.ObjectID).Hex())
	}

	return insertedIDs, err
}
