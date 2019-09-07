package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestLogout(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Logouts", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)
		testUtils.SetupData(db, "sessions", mockSession)

		err := usersController.Logout(mockSession.ID.Hex())

		if err != nil {
			t.Error(err)
		}

		if !testUtils.IsCollectionEmpty(db, "sessions") {
			t.Error("Expected sessions collection to be empty")
		}

	})

	t.Run("No session to logout from", func(t *testing.T) {
		testUtils.ResetCollection(db, "sessions")

		err := usersController.Logout(mockSession.ID.Hex())

		expectedError := "No session with this ID exists"
		if err.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' error but instead got " + err.Error())
		}
	})
}
