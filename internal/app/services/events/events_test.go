package events_test

import (
	"testing"

	testingDb "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/events"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var database = testingDb.MakeDatabaseDriver()
var reps = repositories.MakeRepositories(database)
var service = events.MakeService(reps)

func TestValidate(t *testing.T) {
	t.Run("Should return true for valid event", func(t *testing.T) {
		validEvent := models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "test",
			Description: "test",
			Date:        1583586430,
			Tags:        []string{},
			Type:        "test",
			Sponsors:    []models.Sponsor{},
			Logo:        "test",
			Media:       []string{},
		}
		eventIsValid := !service.Validate(validEvent).HasError()
		if !eventIsValid {
			t.Error("Event shouls be valid")
		}
	})
	t.Run("Should return false for old date", func(t *testing.T) {
		validEvent := models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "test",
			Description: "test",
			Date:        315579661, // 1980/01/01
			Tags:        []string{},
			Type:        "test",
			Sponsors:    []models.Sponsor{},
			Logo:        "test",
			Media:       []string{},
		}
		validation := *service.Validate(validEvent)
		eventIsValid := !validation.HasError()
		if eventIsValid {
			t.Error("Event should be invalid")
		}
		if validation["date"] != events.ErrInvalidDate {
			t.Error("Expected \"" + events.ErrInvalidDate.Error() + "\" but got \"" + validation["date"].Error() + "\"")
		}
	})
	t.Run("Should return false for empty name", func(t *testing.T) {
		validEvent := models.Event{
			ID:          primitive.NewObjectID(),
			Name:        "",
			Description: "test",
			Date:        1583586430,
			Tags:        []string{},
			Type:        "test",
			Sponsors:    []models.Sponsor{},
			Logo:        "test",
			Media:       []string{},
		}
		validation := *service.Validate(validEvent)
		eventIsValid := !validation.HasError()
		if eventIsValid {
			t.Error("Event should be invalid")
		}
		if validation["name"] != events.ErrNameEmpty {
			t.Error("Expected \"" + events.ErrNameEmpty.Error() + "\" but got \"" + validation["name"].Error() + "\"")
		}
	})
}
