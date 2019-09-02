package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/testUtils"
)

func TestLogin(t *testing.T) {
	// Setup
	db, usersController := makeController()

	// Ok login
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

	// Wrong pass
	testUtils.SetupData(db, "users", mockUser)

	sessionID, err = usersController.Login(mockUser.Username, "wrong pass")

	expectedError := "Password not verified"
	if err.Error() != expectedError {
		t.Error("Expected '" + expectedError + "' error but instead got " + err.Error())
	}

	if sessionID != "" {
		t.Error("Expected sessionID to be empty string but instead got " + sessionID)
	}

	// User doesn't exist
	testUtils.ResetCollection(db, "users")

	sessionID, err = usersController.Login("randomuser", "randompass")

	expectedError = "No user found"
	if err.Error() != expectedError {
		t.Error("Expected '" + expectedError + "' error but instead got " + err.Error())
	}

	if sessionID != "" {
		t.Error("Expected sessionID to be empty string but instead got " + sessionID)
	}
}
