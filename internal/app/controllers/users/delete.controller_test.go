package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDelete(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Should delete stored user", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		gotErr := usersController.Delete(mockUser.ID.Hex())

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		if !testUtils.IsCollectionEmpty(db, "users") {
			t.Error("Expected the collection be empty, instead it contains", db.Session.Collections["users"])
		}
	})

	t.Run("Should return invalid ObjectID error", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		gotErr := usersController.Delete("invalid object id")

		expectedError := "Invalid ObjectID"
		if gotErr.Error() != expectedError {
			t.Error("Expected gotError to be '"+expectedError+"' but is", gotErr)
			t.Skip()
		}

		// No such user ObjectID
		testUtils.ResetCollection(db, "users")

		objectID := primitive.NewObjectID()
		gotErr = usersController.Delete(objectID.Hex())

		expectedError = "No document found with this id " + objectID.Hex()
		if gotErr.Error() != expectedError {
			t.Error("Expected gotError to be '"+expectedError+"' but is", gotErr)
		}

	})
}
