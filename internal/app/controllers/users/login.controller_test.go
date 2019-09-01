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
}
