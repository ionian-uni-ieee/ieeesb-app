package testing_test

import (
	"reflect"
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	sponsors "github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories/sponsors/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() (*testingDatabase.DatabaseSession, *sponsors.Repository) {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	sponsorsRepository := sponsors.MakeRepository(database)

	return database, sponsorsRepository
}

var testSponsor1 = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Microsoft",
	Emails: []string{"info@microsoft.com"},
	Phones: []string{"+1 5412871236421"},
	Logo:   "",
}

var testSponsor2 = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Apple",
	Emails: []string{"info@apple.com"},
	Phones: []string{"+1 32462346234523"},
	Logo:   "",
}

var testSponsor3 = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Netflix",
	Emails: []string{"info@netflix.com"},
	Phones: []string{"+1 63242312346124"},
	Logo:   "",
}

func TestFindByID(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1)

	gotSponsor, gotErr := sponsorsRepository.FindByID(testSponsor1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotSponsor == nil {
		t.Error("Expected result to be an sponsor object, got nil instead")
	}

	if gotSponsor != nil && gotSponsor.ID != testSponsor1.ID {
		t.Error("Expected sponsor's id to be", testSponsor1.ID.Hex(), "but is", gotSponsor.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1)

	newName := "New Name"
	gotErr := sponsorsRepository.UpdateByID(testSponsor1.ID.Hex(), map[string]interface{}{
		"Name": newName,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	storedName := sponsorsRepository.Collection.Columns["Name"][0]
	nameChanged := storedName != newName
	if nameChanged {
		t.Error("Expected name to be '"+newName+"', but instead got", storedName)
	}
}

func TestDeleteByID(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1)

	gotErr := sponsorsRepository.DeleteByID(testSponsor1.ID.Hex())

	if gotErr != nil {
		t.Error(gotErr)
	}

	for key, column := range sponsorsRepository.Collection.Columns {
		if len(column) > 0 {
			t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
		}
	}
}

func TestFind(t *testing.T) {
	// Setup
	db, sponsorsRepository := makeRepository()

	t.Run("Should return 2 sponsors", func(t *testing.T) {
		testUtils.SetupData(db, "sponsors", testSponsor1, testSponsor1)

		gotSponsors, gotErr := sponsorsRepository.Find(map[string]interface{}{
			"Name": testSponsor1.Name,
		}, 0, 2)

		if gotErr != nil {
			t.Error(gotErr)
		}

		if len(gotSponsors) != 2 {
			t.Error("Expected length of sponsors got to be 2, instead got", len(gotSponsors))
		}

		if gotSponsors[0].Name != gotSponsors[1].Name {
			t.Error("Expected ID to equal to each other, instead got",
				gotSponsors[0].Name,
				gotSponsors[1].Name)
		}
	})

	t.Run("Should limit the batch to 2 sponsors", func(t *testing.T) {
		testUtils.SetupData(db, "sponsors", testSponsor1, testSponsor2, testSponsor3)

		gotSponsors, gotErr := sponsorsRepository.Find(map[string]interface{}{}, 0, 2)

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		if len(gotSponsors) != 2 {
			t.Error("Expected length of 2 but got", len(gotSponsors))
		}

	})
}

func TestFindOne(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1, testSponsor2)

	gotSponsor, gotErr := sponsorsRepository.FindOne(map[string]interface{}{
		"Name": testSponsor1.Name,
	})

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotSponsor.Name != testSponsor1.Name {
		t.Error("Expected name to equal '" + testSponsor1.Name + "', instead got " + gotSponsor.Name)
	}
}

func TestUpdateMany(t *testing.T) {
	// TODO: Not implemented
}

func TestDeleteMany(t *testing.T) {
	// TODO: Not implemented
}

func TestInsertOne(t *testing.T) {

	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "sponsors")

	gotInsertedID, gotErr := sponsorsRepository.InsertOne(testSponsor1)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedID != testSponsor1.ID.Hex() {
		t.Error("Expected inserted id to be ", testSponsor1.ID.Hex(), "but instead got", gotInsertedID)
	}

	storedSponsor := testUtils.GetInterfaceAtCollectionRow(
		db,
		"sponsors",
		reflect.TypeOf(models.Sponsor{}),
		0,
	).(models.Sponsor)

	if storedSponsor.ID.Hex() != gotInsertedID {
		t.Error("Expected stored user's ID to equal gotInsertedID")
	}
}

func TestInsertMany(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.ResetCollection(db, "sponsors")

	sponsors := []models.Sponsor{
		testSponsor1,
		testSponsor2,
		testSponsor3,
	}

	gotInsertedIDs, gotErr := sponsorsRepository.InsertMany(sponsors)

	if gotErr != nil {
		t.Error(gotErr)
	}

	if gotInsertedIDs[0] != sponsors[0].ID.Hex() ||
		gotInsertedIDs[1] != sponsors[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", sponsors[0].ID.Hex(), sponsors[1].ID.Hex(), "but instead got", gotInsertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Sponsor is duplicate
	testUtils.SetupData(db, "sponsors", testSponsor1)

	gotIsDuplicate := sponsorsRepository.IsDuplicate(testSponsor1.Name)

	if !gotIsDuplicate {
		t.Error("Expected sponsor to be duplicate")
	}

	// Sponsor is not duplicate
	testUtils.ResetCollection(db, "sponsors")

	gotIsDuplicate = sponsorsRepository.IsDuplicate(testSponsor1.Name)

	if gotIsDuplicate {
		t.Error("Expected sponsor to not be duplicate")
	}
}
