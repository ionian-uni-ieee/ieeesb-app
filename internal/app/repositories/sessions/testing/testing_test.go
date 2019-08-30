package sessions_test

import (
	"testing"

	testingDatabase "github.com/ionian-uni-ieee/ieee-webapp/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	sessions "github.com/ionian-uni-ieee/ieee-webapp/internal/app/repositories/sessions/testing"
	"github.com/ionian-uni-ieee/ieee-webapp/pkg/reflections"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func makeRepository() *sessions.Repository {
	// Setup
	database := testingDatabase.MakeDatabaseDriver()
	sessionsRepository := sessions.MakeRepository(database)

	return sessionsRepository
}

// Clears the collection's data
func resetCollection(repository *sessions.Repository) {
	for key, _ := range repository.Collection.Columns {
		repository.Collection.Columns[key] = []interface{}{}
	}
}

// setupData resets the collection and inserts an array of data in it
func setupData(repository *sessions.Repository, data ...models.Session) {
	resetCollection(repository)

	sessionFieldNames, err := reflections.GetFieldNames(&models.Session{})
	if err != nil {
		panic(err)
	}

	for _, item := range data {
		for _, fieldName := range sessionFieldNames {
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
	sessionsRepository := makeRepository()

	// Regular example
	session := models.Session{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(sessionsRepository, session)

	sessionFound, err := sessionsRepository.FindByID(event.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if sessionFound == nil {
		t.Error("Expected result to be an session object, got nil instead")
	}

	if sessionFound != nil && sessionFound.ID != event.ID {
		t.Error("Expected session's id to be", session.ID.Hex(), "but is", eventFound.ID.Hex())
	}
}

func TestUpdateByID(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	session := models.Session{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(sessionsRepository, session)

	err := sessionsRepository.UpdateByID(session.ID.Hex(), map[string]interface{}{
		"Name": "new name",
	})

	if err != nil {
		t.Error(err)
	}

	if name := sessionsRepository.Collection.Columns["Name"][0]; name != "new name" {
		t.Error("Expected session name to be 'new name', but instead got", name)
	}
}

func TestDeleteByID(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	session := models.Session{
		ID:          primitive.NewObjectID(),
		Name:        "name",
		Description: "desc",
		Tags:        []string{"tag1"},
		Type:        "seminar",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	setupData(sessionsRepository, session)

	err := sessionsRepository.DeleteByID(session.ID.Hex())

	if err != nil {
		t.Error(err)
	}

	if ids := sessionsRepository.Collection.Columns["ID"]; len(ids) == 0 {
		t.Error("Expected session id to have length of 0, but instead got", len(ids))
	}
}

func TestFind(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
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
	setupData(sessionsRepository, sessions...)

	sessionsFound, err := sessionsRepository.Find(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if len(sessionsFound) != 2 {
		t.Error("Expected len(sessions) to be 2, instead got", len(sessionsFound))
	}

	if sessionsFound[0].Description != sessionsFound[1].Description {
		t.Error("Expected sessions' description to equal to each other, instead got",
			sessionsFound[0].Description,
			sessionsFound[1].Description)
	}
}

func TestFindOne(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
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
	setupData(sessionsRepository, sessions...)

	sessionFound, err := sessionsRepository.FindOne(map[string]interface{}{
		"Description": "desc3",
	})

	if err != nil {
		t.Error(err)
	}

	if sessionFound.Description != "desc3" {
		t.Error("Expected session description to equal 'desc3', instead got", sessionFound.Description)
	}
}

func TestUpdateMany(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
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
	setupData(sessionsRepository, sessions...)
}

func TestDeleteMany(t *testing.T) {

	sessionsRepository := makeRepository()

	// Regular example
	sessions := []models.Session{
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name2",
			Description: "desc3",
			Tags:        []string{"tag2"},
			Type:        "workshop",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
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
	setupData(sessionsRepository, sessions...)
}

func TestInsertOne(t *testing.T) {

	sessionsRepository := makeRepository()

	// Regular example
	resetCollection(sessionsRepository)

	newSession := models.Event{
		ID:          primitive.NewObjectID(),
		Name:        "name3",
		Description: "desc3",
		Tags:        []string{"tag3"},
		Type:        "workshop",
		Sponsors:    []models.Sponsor{},
		Logo:        models.MediaMeta{},
		Media:       []models.MediaMeta{},
	}
	insertedID, err := sessionsRepository.InsertOne(newSession)

	if err != nil {
		t.Error(err)
	}

	if insertedID != newSession.ID.Hex() {
		t.Error("Expected inserted id to be ", newSession.ID.Hex(), "but instead got", insertedID)
	}
}

func TestInsertMany(t *testing.T) {
	sessionsRepository := makeRepository()

	// Regular example
	resetCollection(sessionsRepository)

	newSessions := []models.Event{
		models.Session{
			ID:          primitive.NewObjectID(),
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1"},
			Type:        "seminar",
			Sponsors:    []models.Sponsor{},
			Logo:        models.MediaMeta{},
			Media:       []models.MediaMeta{},
		},
		models.Session{
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

	insertedIDs, err := sessionsRepository.InsertMany(newSessions)

	if err != nil {
		t.Error(err)
	}

	if insertedIDs[0] != newSessions[0].ID.Hex() ||
		insertedIDs[1] != newSessions[1].ID.Hex() {
		t.Error("Expected inserted ids to be ", newSessions[0].ID.Hex(), newEvents[1].ID.Hex(), "but instead got", insertedIDs)
	}
}

func TestIsDuplicate(t *testing.T) {
	sessionsRepository := makeRepository()

	// Name is duplicate
	sessions := []models.Session{
		models.Session{
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
	setupData(sessionsRepository, sessions...)

	isDuplicate := sessionsRepository.IsDuplicate("name2")

	if !isDuplicate {
		t.Error("Expected name to be duplicate")
	}
}
