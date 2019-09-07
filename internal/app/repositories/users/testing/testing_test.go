package testing_test

import (
	"reflect"
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	users "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/users/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
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

	t.Run("Finds user", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		userFound, err := usersRepository.FindByID(testUser1.ID.Hex())

		if err != nil {
			t.Error(err)
		}

		if userFound == nil {
			t.Error("Expected result to be an user object, got nil instead")
		}

		if userFound != nil && userFound.ID != testUser1.ID {
			t.Error("Expected user's id to be", testUser1.ID.Hex(), "but is", userFound.ID.Hex())
		}
	})
}

func TestUpdateByID(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Updates user", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		newUsername := "newusername"
		err := usersRepository.UpdateByID(testUser1.ID.Hex(), map[string]interface{}{
			"Username": newUsername,
		})

		if err != nil {
			t.Error(err)
		}

		usernameStored := usersRepository.Collection.Columns["Username"][0]
		usernameChanged := usernameStored != newUsername
		if usernameChanged {
			t.Error("Expected username to be '"+newUsername+"', but instead got", usernameStored)
		}
	})
}

func TestDeleteByID(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Deletes user", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		err := usersRepository.DeleteByID(testUser1.ID.Hex())

		if err != nil {
			t.Error(err)
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

	t.Run("Finds users", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1, testUser1)

		usersFound, err := usersRepository.Find(map[string]interface{}{
			"Email": testUser1.Email,
		})

		if err != nil {
			t.Error(err)
		}

		if len(usersFound) != 2 {
			t.Error("Expected len(users) to be 2, instead got", len(usersFound))
		}

		if usersFound[0].Username != usersFound[1].Username {
			t.Error("Expected username to equal to each other, instead got",
				usersFound[0].Username,
				usersFound[1].Username)
		}
	})
}

func TestFindOne(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Find a user", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1, testUser2)

		userFound, err := usersRepository.FindOne(map[string]interface{}{
			"Username": testUser1.Username,
		})

		if err != nil {
			t.Error(err)
		}

		if userFound.Username != testUser1.Username {
			t.Error("Expected username to equal 'username2', instead got", userFound.Username)
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

	t.Run("Inserts a user", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		insertedID, err := usersRepository.InsertOne(testUser1)

		if err != nil {
			t.Error(err)
		}

		if insertedID != testUser1.ID.Hex() {
			t.Error("Expected inserted id to be ", testUser1.ID.Hex(), "but instead got", insertedID)
		}

		storedUser := testUtils.GetInterfaceAtCollectionRow(
			db,
			"users",
			reflect.TypeOf(models.User{}),
			0,
		).(models.User)

		if storedUser.ID.Hex() != insertedID {
			t.Error("Expected stored user's ID to equal insertedID")
		}
	})
}

func TestInsertMany(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Inserts many users", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		users := []models.User{
			testUser1,
			testUser2,
			testUser3,
		}

		insertedIDs, err := usersRepository.InsertMany(users)

		if err != nil {
			t.Error(err)
		}

		if insertedIDs[0] != users[0].ID.Hex() ||
			insertedIDs[1] != users[1].ID.Hex() {
			t.Error("Expected inserted ids to be ", users[0].ID.Hex(), users[1].ID.Hex(), "but instead got", insertedIDs)
		}
	})
}

func TestIsDuplicate(t *testing.T) {
	db, usersRepository := makeRepository()

	t.Run("Is duplicate user", func(t *testing.T) {
		testUtils.SetupData(db, "users", testUser1)

		isDuplicate := usersRepository.IsDuplicate(testUser1.Email, testUser1.Username, testUser1.Fullname)

		if !isDuplicate {
			t.Error("Expected user to be duplicate")
		}
	})

	t.Run("Is not a duplicate user", func(t *testing.T) {
		testUtils.ResetCollection(db, "users")

		isDuplicate := usersRepository.IsDuplicate(testUser1.Email, testUser1.Username, testUser1.Fullname)

		if isDuplicate {
			t.Error("Expected user to not be duplicate")
		}
	})
}
