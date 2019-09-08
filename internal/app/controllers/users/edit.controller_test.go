package users_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestEdit(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Edits user", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser)

		newFullname := "New Fullname"
		updateMap := map[string]interface{}{
			"Fullname": "New Fullname",
		}
		gotErr := usersController.Edit(mockUser.ID.Hex(), updateMap)

		if gotErr != nil {
			t.Error(gotErr)
		}

		storedUser := testUtils.GetInterfaceAtCollectionRow(
			db,
			"users",
			reflect.TypeOf(models.User{}),
			0,
		).(models.User)

		if storedUser.Fullname != newFullname {
			t.Error("Expected fullname to have been changed, but it's " + storedUser.Fullname)
		}
	})
}
