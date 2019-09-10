package sponsors_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestEdit(t *testing.T) {
	// Setup
	db, sponsorsController := makeController()

	t.Run("Should change stored sponsor's name", func(t *testing.T) {
		testUtils.SetupData(db, "sponsors", mockSponsor)

		update := map[string]interface{}{
			"Name": "new sponsor",
		}
		gotErr := sponsorsController.Edit(mockSponsor.ID.Hex(), update)

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

		didNameChange := storedSponsor.Name == update["Name"]
		if !didNameChange {
			t.Error("Expected name to be '" + update["Name"].(string) + "' but instead its '" + storedSponsor.Name)
		}
	})
}
