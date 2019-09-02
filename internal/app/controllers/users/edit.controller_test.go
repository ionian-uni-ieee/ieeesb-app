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

	// Ok
	testUtils.SetupData(db, "users", mockUser)

	newFullname := "New Fullname"
	updateMap := map[string]interface{}{
		"Fullname": "New Fullname",
	}
	err := usersController.Edit(mockUser.ID.Hex(), updateMap)

	if err != nil {
		t.Error(err)
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
}
