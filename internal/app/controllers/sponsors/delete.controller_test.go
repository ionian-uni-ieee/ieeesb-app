package sponsors_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/models"
	"github.com/ionian-uni-ieee/ieee-webapp/internal/app/testUtils"
)

func TestDelete(t *testing.T) {
	// Setup
	db, sponsorsController := makeController()

	t.Run("Should delete sponsor", func(t *testing.T) {
		testUtils.SetupData(db, "sponsors", mockSponsor)

		gotErr := sponsorsController.Delete(mockSponsor.ID.Hex())

		if gotErr != nil {
			t.Error(gotErr)
			t.SkipNow()
		}

		storedSponsor := testUtils.GetInterfaceAtCollectionRow(
			db,
			"sponsors",
			reflect.TypeOf(models.Sponsor{}),
			0,
		).(models.Sponsor)

		if !storedSponsor.ID.IsZero() {
			t.Error("Expected sponsor to have been removed from the repository")
		}
	})
}
