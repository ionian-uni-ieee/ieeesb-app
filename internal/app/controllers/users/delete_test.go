package users_test

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDelete(t *testing.T) {
	// Setup
	db, usersController := makeController()

	// Deletable user
	setupData(db, "users", mockUser)

	err := usersController.Delete(mockUser.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if !isCollectionEmpty(db, "users") {
		t.Error("Expected the collection be empty, instead it contains", db.Session.Collections["users"])
	}

	// Invalid ObjectID
	resetCollection(db, "users")

	err = usersController.Delete("invalid object id")

	if err.Error() != "Invalid ObjectID" {
		t.Error("Expected error to be 'Invalid ObjectID' but is", err)
	}

	// No such user ObjectID
	resetCollection(db, "users")

	objectID := primitive.NewObjectID()
	err = usersController.Delete(objectID.Hex())

	if err.Error() != ("No document found with this id " + objectID.Hex()) {
		t.Error("Expected error to be 'No document found with this id "+objectID.Hex()+"' but is", err)
	}
}
