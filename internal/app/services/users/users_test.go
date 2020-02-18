package users_test

import (
	"testing"

	testingDb "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/users"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var database = testingDb.MakeDatabaseDriver()
var reps = repositories.MakeRepositories(database)
var service = users.MakeService(reps)

func TestValidate(t *testing.T) {
	t.Run("Should return true for valid user", func(t *testing.T) {
		validUser := models.User{
			ID:          primitive.NewObjectID(),
			Email:       "test@test.com",
			Fullname:    "test",
			Password:    "test",
			Permissions: models.Permissions{},
			Username:    "test",
		}
		userIsValid := !service.Validate(validUser).HasError()
		if !userIsValid {
			t.Error("User should be valid")
		}
	})
	t.Run("Should return false for wrong email form", func(t *testing.T) {
		validUser := models.User{
			ID:          primitive.NewObjectID(),
			Email:       "wrongemailform",
			Fullname:    "test",
			Password:    "test",
			Permissions: models.Permissions{},
			Username:    "test",
		}

		validation := *service.Validate(validUser)
		userIsValid := !validation.HasError()
		if userIsValid {
			t.Error("User should be invalid")
		}
		if validation["email"] != users.ErrInvalidEmail {
			t.Error("Expected \"" + users.ErrInvalidEmail.Error() + "\" but got \"" + validation["email"].Error() + "\"")
		}
	})
}
