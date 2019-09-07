package users_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestChangePassword(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Correct password", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		newPassword := "new pass"

		err := usersController.ChangePassword(mockUser.ID.Hex(), mockUserPass, newPassword)

		if err != nil {
			t.Error(err)
		}

		firstUser := testUtils.GetInterfaceAtCollectionRow(
			db,
			"users",
			reflect.TypeOf(models.User{}),
			0).(models.User)

		if firstUser.Password == mockUser.Password {
			t.Error("Expected password to have been changed, but it's the same with the old one")
		}
	})

	t.Run("Wrong password", func(t *testing.T) {

		testUtils.SetupData(db, "users", mockUser)

		err := usersController.ChangePassword(mockUser.ID.Hex(), "wrong pass", "new password")

		if err.Error() != "Old password is not valid" {
			t.Error("Expected error to be 'Old password is not valid' but instead is ", err)
		}

		firstUser := testUtils.GetInterfaceAtCollectionRow(
			db,
			"users",
			reflect.TypeOf(models.User{}),
			0).(models.User)

		if firstUser.Password != mockUser.Password {
			t.Error("Expected password to not have been changed")
		}
	})
}
