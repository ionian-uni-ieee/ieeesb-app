package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProfile(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Gets profile", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)
		testUtils.SetupData(db, "sessions", mockSession)

		user, err := usersController.Profile(mockSession.ID.Hex())

		if err != nil {
			t.Error(err)
		}

		if user.ID != mockUser.ID {
			t.Error("Expected profile to have ID " + mockUser.ID.Hex() + ", instead it has " + user.ID.Hex())
		}
	})

	// No such session
	t.Run("No session", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")
		testUtils.ResetCollection(db, "sessions")

		_, err := usersController.Profile(primitive.NewObjectID().Hex())

		expectedError := "No session with this ID exists"
		if err.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' error but instead got " + err.Error())
		}
	})
}
