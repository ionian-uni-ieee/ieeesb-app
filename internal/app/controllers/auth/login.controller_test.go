package auth_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestLogin(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Should store a new session", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		gotSessionID, gotErr := usersController.Login(mockUser.Username, mockUserPass)

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		if gotSessionID == "" {
			t.Error("Expected gotSessionID to be a non empty string")
			t.SkipNow()
		}

		storedSession := testUtils.GetInterfaceAtCollectionRow(
			db,
			"sessions",
			reflect.TypeOf(models.Session{}),
			0,
		).(models.Session)

		if storedSession.ID.Hex() != gotSessionID {
			t.Error("Expected gotSessionID to match the stored session's ID, instead got " + storedSession.ID.Hex())
		}
	})

	t.Run("Should return wrong password error", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		gotSessionID, gotErr := usersController.Login(mockUser.Username, "wrong pass")

		expectedError := "Password not verified"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but instead got " + gotErr.Error())
			t.SkipNow()
		}

		if gotSessionID != "" {
			t.Error("Expected gotSessionID to be empty string but instead got " + gotSessionID)
		}
	})

	t.Run("Should return user not found error", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		gotSessionID, gotErr := usersController.Login("randomuser", "randompass")

		expectedError := "No user found"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but instead got " + gotErr.Error())
			t.SkipNow()
		}

		if gotSessionID != "" {
			t.Error("Expected gotSessionID to be empty string but instead got " + gotSessionID)
		}
	})
}
