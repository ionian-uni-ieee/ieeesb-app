package users_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestGetUsers(t *testing.T) {
	// Setup
	db, usersController := makeController()

	t.Run("Should return an array of 2 users", func(t *testing.T) {
		testUtils.SetupData(db, "users", mockUser, mockUser2)

		gotUsers, err := usersController.GetUsers(0, 2)

		if err != nil {
			t.Error(err)
			t.SkipNow()
		}

		if len(gotUsers) != 2 {
			t.Error("Expected users array with length of 2, but got", len(gotUsers))
			t.SkipNow()
		}

		gotValidData := gotUsers[0].ID == mockUser.ID ||
			gotUsers[1].ID == mockUser2.ID
		if !gotValidData {
			t.Error("Expected returned data to have the same IDs with stored data")
		}
	})
}
