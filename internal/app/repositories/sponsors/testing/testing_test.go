package sponsors_test

import (
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	sponsors "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sponsors/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() *sponsors.Repository {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	sponsorsRepository := sponsors.MakeRepository(database)

	return sponsorsRepository
}

// Clears the collection's data
func resetCollection(repository *sponsors.Repository) {
	for key, _ := range repository.Collection.Columns {
		repository.Collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(repository *sponsors.Repository, data ...models.Sponsor) {
	resetCollection(repository)

	sponsorFieldNames, err := reflections.GetFieldNames(&models.Sponsor{})
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range sponsorFieldNames {
			fieldValue, err := reflections.GetField(&item, fieldName)
			if err != nil {
				panic(err)
			}

			repository.Collection.Columns[fieldName] = append(
				repository.Collection.Columns[fieldName],
				fieldValue)
		}
	}
}

func TestFindByID(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	sponsor := models.Sponsor{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(sponsorsRepository, sponsor)

	sponsorFound, err := sponsorsRepository.FindByID(event.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if sponsorFound == nil {
		t.Error("Expected result to be an sponsor object, got nil instead")
	}

	if sponsorFound != nil && sponsorFound.ID != event.ID {
		t.Error("Expected sponsor's id to be", sponsor.ID.Hex(), "but is", eventFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	sponsor := models.Sponsor{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(sponsorsRepository, sponsor)

	err := sponsorsRepository.UpdateByID(sponsor.ID.Hex(), map[string]interface{}{
		"Name": "new name",
	})

	if err != nil {
		t.Error(err)
	}

	if name := sponsorsRepository.Collection.Columns["Name"][0]; name != "new name" {
		t.Error("Expected sponsor name to be 'new name', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	sponsor := models.Sponsor{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(sponsorsRepository, sponsor)

	err := sponsorsRepository.DeleteByID(sponsor.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ids := sponsorsRepository.Collection.Columns["ID"]; len(ids) == 0 {
		t.Error("Expected sponsor id to have length of 0, but instead got", len(ids))
	}
}

func TestFind(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	sponsors := []models.Sponsor{
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(sponsorsRepository, sponsors...)

	sponsorsFound, err := sponsorsRepository.Find(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if len(sponsorsFound) != 2 {
		t.Error("Expected len(sponsors) to be 2, instead got", len(sponsorsFound))
	}

	if sponsorsFound[0].Description != sponsorsFound[1].Description {
		t.Error("Expected sponsors' description to equal to each other, instead got",
			sponsorsFound[0].Description,
			sponsorsFound[1].Description)
	}
}

func TestFindOne(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	sponsors := []models.Sponsor{
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(sponsorsRepository, sponsors...)

	sponsorFound, err := sponsorsRepository.FindOne(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if sponsorFound.Description != "desc3" {
		t.Error("Expected sponsor description to equal 'desc3', instead got", sponsorFound.Description)
	}
}

func TestUpdateMany(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	sponsors := []models.Sponsor{
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(sponsorsRepository, sponsors...)
}

func TestDeleteMany(t *testing.T) {

	sponsorsRepository := makeRepository()

	// Regular example
	sponsors := []models.Sponsor{
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name3",
			Description: "desc3",
			Tags:        []string{"tag3"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(sponsorsRepository, sponsors...)
}

func TestInsertOne(t *testing.T) {

	sponsorsRepository := makeRepository()

	// Regular example
	resetCollection(sponsorsRepository)

	newSponsor := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name3",
		Description: "desc3",
		Tags:        []string{"tag3"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	insertedID, err := sponsorsRepository.InsertOne(newSponsor)

	if err != nil {
		t.Error(err)
	}

	if insertedID != newSponsor.ID.Hex() {
		t.Error("Expected inserted id to be ", newSponsor.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Regular example
	resetCollection(sponsorsRepository)

	newSponsors := []models.Event{
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}

	insertedIDs, err := sponsorsRepository.InsertMany(newSponsors)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newSponsors[0].ID.Hex() ||
		insertedIDs[1] != newSponsors[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newSponsors[0].ID.Hex(), newEvents[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	sponsorsRepository := makeRepository()

	// Name is duplicate
	sponsors := []models.Sponsor{
		models.Sponsor{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
	}
	setupData(sponsorsRepository, sponsors...)

	isDuplicate := sponsorsRepository.IsDuplicate("name2")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
