package testing_test

import (
	"reflect"
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	users "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/users/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() (*testingDatabase.DatabaseSession, *users.Repository) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	usersRepository := users.MakeRepository(database)

	return database, usersRepository
}

var testUser1 = models.User{
	ID:          primitive.NewObjectID(),
	Username:    "joe",
	Password:    "joepassftw",
	Email:       "joe@mail.com",
	Fullname:    "joe jordinson",
	Permissions: models.Permissions{},
}

var testUser2 = models.User{
	ID:          primitive.NewObjectID(),
	Username:    "johndoe",
	Password:    "hmm",
	Email:       "johndoe@mail.com",
	Fullname:    "John Doe",
	Permissions: models.Permissions{},
}

var testUser3 = models.User{
	ID:          primitive.NewObjectID(),
	Username:    "smith",
	Password:    "1983billsmith",
	Email:       "billsmith@mail.com",
	Fullname:    "Bill Smith",
	Permissions: models.Permissions{},
}

func TestFindByID(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should return a user", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		gotUser, err := usersRepository.FindByID(testUser1.ID.Hex())

		if err != nil {
			t.Error(err)
		}

		if gotUser == nil {
			t.Error("Expected result to be an user object, got nil instead")
		}

		if gotUser != nil && gotUser.ID != testUser1.ID {
			t.Error("Expected user's id to be", testUser1.ID.Hex(), "but is", gotUser.ID.Hex())
		}
	})
}

func TestUpdateByID(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should update a stored user's data", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		newUsername := "newusername"

		gotErr := usersRepository.UpdateByID(testUser1.ID.Hex(), map[string]interface{}{
			"Username": newUsername,
		})

		if gotErr != nil {
			t.Error(gotErr)
		}

		storedUsername := usersRepository.Collection.Columns["Username"][0]

		usernameChanged := storedUsername != newUsername
		if usernameChanged {
			t.Error("Expected username to be '"+newUsername+"', but instead got", storedUsername)
		}
	})
}

func TestDeleteByID(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should delete a user from the database", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		gotErr := usersRepository.DeleteByID(testUser1.ID.Hex())

		if gotErr != nil {
			t.Error(gotErr)
		}

		for key, column := range usersRepository.Collection.Columns {
			if len(column) > 0 {
				t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
			}
		}
	})
}

func TestFind(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should return 2 users", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1, testUser1)

		gotUsers, gotErr := usersRepository.Find(map[string]interface{}{
			"Email": testUser1.Email,
		}, 0, 2)

		if gotErr != nil {
			t.Error(gotErr)
		}

		if len(gotUsers) != 2 {
			t.Error("Expected length of users got to be 2, instead got", len(gotUsers))
		}

		if gotUsers[0].Username != gotUsers[1].Username {
			t.Error("Expected username to equal to each other, instead got",
				gotUsers[0].Username,
				gotUsers[1].Username)
		}
	})

	t.Run("Should limit the batch to 2 users", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1, testUser2, testUser3)

		gotUsers, gotErr := usersRepository.Find(map[string]interface{}{}, 0, 2)

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		if len(gotUsers) != 2 {
			t.Error("Expected length of 2 but got", len(gotUsers))
		}

	})
}

func TestFindOne(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should return a user with same username as the stored one", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1, testUser2)

		gotUser, gotErr := usersRepository.FindOne(map[string]interface{}{
			"Username": testUser1.Username,
		})

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotUser.Username != testUser1.Username {
			t.Error("Expected username to equal 'username2', instead got", gotUser.Username)
		}
	})
}

func TestUpdateMany(t *testing.T) {
	// TODO: Not implemented
}

func TestDeleteMany(t *testing.T) {
	// TODO: Not implemented
}

func TestInsertOne(t *testing.T) {

	db, usersRepository := makeRepository()

	t.Run("Should insert a new user into database", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		gotInsertedID, gotErr := usersRepository.InsertOne(testUser1)

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotInsertedID != testUser1.ID.Hex() {
			t.Error("Expected inserted id to be ", testUser1.ID.Hex(), "but instead got", gotInsertedID)
		}

		storedUser := testUtils.GetInterfaceAtCollectionRow(
			db,
			"users",
			reflect.TypeOf(models.User{}),
			0,
		).(models.User)

		if storedUser.ID.Hex() != gotInsertedID {
			t.Error("Expected stored user's ID to equal insertedID")
		}
	})
}

func TestInsertMany(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should store many users into database", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		users := []models.User{
			testUser1,
			testUser2,
			testUser3,
		}

		gotInsertedIDs, gotErr := usersRepository.InsertMany(users)

		if gotErr != nil {
			t.Error(gotErr)
		}

		if gotInsertedIDs[0] != users[0].ID.Hex() ||
			gotInsertedIDs[1] != users[1].ID.Hex() {
			t.Error("Expected inserted ids to be ", users[0].ID.Hex(), users[1].ID.Hex(), "but instead got", gotInsertedIDs)
		}
	})
}

func TestIsDuplicate(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Should return that user is duplicate", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		gotIsDuplicate := usersRepository.IsDuplicate(testUser1.Email, testUser1.Username, testUser1.Fullname)

		if !gotIsDuplicate {
			t.Error("Expected user to be duplicate")
		}
	})

	t.Run("Should return that user is not duplicate", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		gotIsDuplicate := usersRepository.IsDuplicate(testUser1.Email, testUser1.Username, testUser1.Fullname)

		if gotIsDuplicate {
			t.Error("Expected user to not be duplicate")
		}
	})
}
