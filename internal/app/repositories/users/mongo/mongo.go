package mongo

import (
	"context"
	"errors"
	"reflect"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongod "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database   database.Driver
	collection *mongod.Collection
}

func MakeRepository(database database.Driver) *Repository {
	collection := database.GetCollection("users").(*mongod.Collection)
	return &Repository{database, collection}
}

func (r *Repository) FindByID(userID string) (*models.User, error) {
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

func (r *Repository) UpdateByID(userID string, update interface{}) error {
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

	updateBSON, err := reflections.ConvertFieldNamesToTagNames(
		update.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(context.Background(), &bson.M{"_id": userObjectID}, &bson.M{"$set": updateBSON})

	return err
}

func (r *Repository) DeleteByID(userID string) error {
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

func (r *Repository) Find(filter interface{}) ([]models.User, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	result, err := r.collection.Find(context.Background(), filterBSON)
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

func (r *Repository) FindOne(filter interface{}) (*models.User, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	result := r.collection.FindOne(context.Background(), filterBSON)

	user := &models.User{}

	err = result.Decode(user)

	return user, err
}

func (r *Repository) UpdateMany(filter interface{}, update interface{}) ([]string, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	updateBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return nil, err
	}

	result, err := r.collection.UpdateMany(context.Background(), filterBSON, updateBSON)

	return result.UpsertedID.([]string), err
}

func (r *Repository) DeleteMany(filter interface{}) (int64, error) {
	filterBSON, err := reflections.ConvertFieldNamesToTagNames(
		filter.(map[string]interface{}),
		reflect.TypeOf(models.User{}),
		"bson",
	)

	if err != nil {
		return -1, err
	}

	result, err := r.collection.DeleteMany(context.Background(), filterBSON)

	return result.DeletedCount, err
}

func (r *Repository) InsertOne(document models.User) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), document)

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r *Repository) InsertMany(documents []models.User) ([]string, error) {
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

func (r *Repository) IsDuplicate(email string, username string, fullname string) bool {
	sameKeysFilter := &bson.M{
		"$or": bson.A{
			bson.M{"email": email},
			bson.M{"username": username},
			bson.M{"fullname": fullname},
		}}

	userFound := models.User{}
	r.collection.FindOne(context.Background(), sameKeysFilter).Decode(&userFound)

	return !userFound.ID.IsZero()
}
