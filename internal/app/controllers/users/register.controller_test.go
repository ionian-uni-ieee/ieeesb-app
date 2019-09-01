package users_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/controllers/testUtils"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
)

func TestRegister(t *testing.T) {
	// Setup
	db, usersController := makeController()

	// Register user
	testUtils.SetupData(db, "users")

	userID, err := usersController.Register("john", "johnpass", "john@mail.com", "john doe")

	if err != nil {
		t.Error(err)
	}

	firstUser := testUtils.GetInterfaceAtCollectionRow(
		db,
		"users",
		reflect.TypeOf(models.User{}),
		0,
	).(models.User)

	if firstUsersID := firstUser.ID.Hex(); firstUsersID != userID {
		t.Error("Expected user's ID at row 0 to be equal to " + userID + " but instead it's equal to " + firstUsersID)
	}

	// Register a duplicate user
	testUtils.SetupData(db, "users", mockUser)

	userID, err = usersController.Register(mockUser.Username, mockUserPass, mockUser.Email, mockUser.Fullname)

	if err != nil {
		t.Error(err)
	}

	if testUtils.GetCollectionLength(db, "users") == 2 {
		t.Error("Expected users collection length to be 1, instead the new user got registered")
	}
}
