package users_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestLogin(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Correct Login", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		sessionID, err := usersController.Login(mockUser.Username, mockUserPass)

		if err != nil {
			t.Error(err)
		}

		if sessionID == "" {
			t.Error("Expected sessionID to be a non empty string")
		}

		storedSession := testUtils.GetInterfaceAtCollectionRow(
			db,
			"sessions",
			reflect.TypeOf(models.Session{}),
			0,
		).(models.Session)

		if storedSession.ID.Hex() != sessionID {
			t.Error("Expected sessionID to match the stored session's ID, instead got " + storedSession.ID.Hex())
		}
	})

	t.Run("Wrong password", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		sessionID, err := usersController.Login(mockUser.Username, "wrong pass")

		expectedError := "Password not verified"
		if err.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' error but instead got " + err.Error())
		}

		if sessionID != "" {
			t.Error("Expected sessionID to be empty string but instead got " + sessionID)
		}
	})

	t.Run("User doesn't exist", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		sessionID, err := usersController.Login("randomuser", "randompass")

		expectedError := "No user found"
		if err.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' error but instead got " + err.Error())
		}

		if sessionID != "" {
			t.Error("Expected sessionID to be empty string but instead got " + sessionID)
		}
	})
}
