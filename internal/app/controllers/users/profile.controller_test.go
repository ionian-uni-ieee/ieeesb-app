package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProfile(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Should return user's profile", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)
		testUtils.SetupData(db, "sessions", mockSession)

		gotUser, gotErr := usersController.Profile(mockSession.ID.Hex())

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		if gotUser.ID != mockUser.ID {
			t.Error("Expected profile to have ID " + mockUser.ID.Hex() + ", instead it has " + gotUser.ID.Hex())
		}
	})

	t.Run("Should return no such session error", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")
		testUtils.ResetCollection(db, "sessions")

		_, gotErr := usersController.Profile(primitive.NewObjectID().Hex())

		expectedError := "No session with this ID exists"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but instead got " + gotErr.Error())
		}
	})
}
