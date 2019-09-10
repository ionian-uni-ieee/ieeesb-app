package users_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestRegister(t *testing.T) {
	// Setup
	db, usersController := makeController()

	// Register user
	t.Run("Should store a new user", func(t *testing.T) {
		testUtils.SetupData(db, "users")

		gotUserID, gotErr := usersController.Register("john", "johnpass", "john@mail.com", "john doe")

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		firstUser := testUtils.GetInterfaceAtCollectionRow(
			db,
			"users",
			reflect.TypeOf(models.User{}),
			0,
		).(models.User)

		firstUserIDEqualsUserID := firstUser.ID.Hex() != gotUserID
		if firstUserIDEqualsUserID {
			t.Error("Expected user's ID at row 0 to be equal to " + gotUserID + " but instead it's equal to " + firstUser.ID.Hex())
		}
	})

	t.Run("Should return user is duplicate error", func(t *testing.T) {

		testUtils.SetupData(db, "users", mockUser)

		_, gotErr := usersController.Register(mockUser.Username, mockUserPass, mockUser.Email, mockUser.Fullname)

		expectedError := "A user with that username, fullname or email already exists"
		if gotErr.Error() != expectedError {
			t.Error("Expected '" + expectedError + "' gotError but instead got " + gotErr.Error())
			t.SkipNow()
		}

		if testUtils.GetCollectionLength(db, "users") == 2 {
			t.Error("Expected users collection length to be 1, instead the new user got registered")
		}
	})
}
