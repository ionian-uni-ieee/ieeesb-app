package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDelete(t *testing.T) {
	// Setup
	db, usersController := makeController()

	// Deletable user
	testUtils.SetupData(db, "users", mockUser)

	err := usersController.Delete(mockUser.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if !testUtils.IsCollectionEmpty(db, "users") {
		t.Error("Expected the collection be empty, instead it contains", db.Session.Collections["users"])
	}

	// Invalid ObjectID
	testUtils.ResetCollection(db, "users")

	err = usersController.Delete("invalid object id")

	expectedError := "Invalid ObjectID"
	if err.Error() != expectedError {
		t.Error("Expected error to be '"+expectedError+"' but is", err)
	}

	// No such user ObjectID
	testUtils.ResetCollection(db, "users")

	objectID := primitive.NewObjectID()
	err = usersController.Delete(objectID.Hex())

	expectedError = "No document found with this id " + objectID.Hex()
	if err.Error() != expectedError {
		t.Error("Expected error to be '"+expectedError+"' but is", err)
	}
}
