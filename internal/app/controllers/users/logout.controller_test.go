package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestLogout(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Should delete stored session", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)
		testUtils.SetupData(db, "sessions", mockSession)

		gotErr := usersController.Logout(mockSession.ID.Hex())

		if gotErr != nil {
			t.Error(gotErr)
			t.Skip()
		}

		if !testUtils.IsCollectionEmpty(db, "sessions") {
			t.Error("Expected sessions collection to be empty")
		}

	})

	t.Run("Should return no such session error", func(t *testing.T) {
		testUtils.ResetCollection(db, "sessions")

		gotErr := usersController.Logout(mockSession.ID.Hex())

		expectedError := "No session with this ID exists"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but instead got " + gotErr.Error())
		}
	})
}
