package testing_test

import (
	"reflect"
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	sponsors "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sponsors/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
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
	Logo:   models.MediaMeta{},
}

var testSponsor2 = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Apple",
	Emails: []string{"info@apple.com"},
	Phones: []string{"+1 32462346234523"},
	Logo:   models.MediaMeta{},
}

var testSponsor3 = models.Sponsor{
	ID:     primitive.NewObjectID(),
	Name:   "Netflix",
	Emails: []string{"info@netflix.com"},
	Phones: []string{"+1 63242312346124"},
	Logo:   models.MediaMeta{},
}

func TestFindByID(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1)

	sponsorFound, err := sponsorsRepository.FindByID(testSponsor1.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if sponsorFound == nil {
		t.Error("Expected result to be an sponsor object, got nil instead")
	}

	if sponsorFound != nil && sponsorFound.ID != testSponsor1.ID {
		t.Error("Expected sponsor's id to be", testSponsor1.ID.Hex(), "but is", sponsorFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1)

	newName := "New Name"
	err := sponsorsRepository.UpdateByID(testSponsor1.ID.Hex(), map[string]interface{}{
		"Name": newName,
	})

	if err != nil {
		t.Error(err)
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

	err := sponsorsRepository.DeleteByID(testSponsor1.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	for key, column := range sponsorsRepository.Collection.Columns {
		if len(column) > 0 {
			t.Error("Expected column", key, "to have length of 0, but instead got", len(column))
		}
	}
}

func TestFind(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1, testSponsor1)

	sponsorsFound, err := sponsorsRepository.Find(map[string]interface{}{
		"Name": testSponsor1.Name,
	})

	if err != nil {
		t.Error(err)
	}

	if len(sponsorsFound) != 2 {
		t.Error("Expected len(sponsors) to be 2, instead got", len(sponsorsFound))
	}

	if sponsorsFound[0].Name != sponsorsFound[1].Name {
		t.Error("Expected sponsorname to equal to each other, instead got",
			sponsorsFound[0].Name,
			sponsorsFound[1].Name)
	}
}

func TestFindOne(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Regular example
	testUtils.SetupData(db, "sponsors", testSponsor1, testSponsor2)

	sponsorFound, err := sponsorsRepository.FindOne(map[string]interface{}{
		"Name": testSponsor1.Name,
	})

	if err != nil {
		t.Error(err)
	}

	if sponsorFound.Name != testSponsor1.Name {
		t.Error("Expected name to equal '" + testSponsor1.Name + "', instead got " + sponsorFound.Name)
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

	insertedID, err := sponsorsRepository.InsertOne(testSponsor1)

	if err != nil {
		t.Error(err)
	}

	if insertedID != testSponsor1.ID.Hex() {
		t.Error("Expected inserted id to be ", testSponsor1.ID.Hex(), "but instead got", insertedID)
	}

	storedSponsor := testUtils.GetInterfaceAtCollectionRow(
		db,
		"sponsors",
		reflect.TypeOf(models.Sponsor{}),
		0,
	).(models.Sponsor)

	if storedSponsor.ID.Hex() != insertedID {
		t.Error("Expected stored user's ID to equal insertedID")
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

	insertedIDs, err := sponsorsRepository.InsertMany(sponsors)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != sponsors[0].ID.Hex() ||
		insertedIDs[1] != sponsors[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", sponsors[0].ID.Hex(), sponsors[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	db, sponsorsRepository := makeRepository()

	// Sponsor is duplicate
	testUtils.SetupData(db, "sponsors", testSponsor1)

	isDuplicate := sponsorsRepository.IsDuplicate(testSponsor1.Name)

	if !isDuplicate {
		t.Error("Expected sponsor to be duplicate")
	}

	// Sponsor is not duplicate
	testUtils.ResetCollection(db, "sponsors")

	isDuplicate = sponsorsRepository.IsDuplicate(testSponsor1.Name)

	if isDuplicate {
		t.Error("Expected sponsor to not be duplicate")
	}
}
