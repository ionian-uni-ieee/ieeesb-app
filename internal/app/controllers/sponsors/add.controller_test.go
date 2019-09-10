package sponsors_test

import (
	"reflect"
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestAdd(t *testing.T) {
	//Setup
	db, sponsorsController := makeController()

	t.Run("Should add a new sponsor", func(t *testing.T) {
		testUtils.ResetCollection(db, "sponsors")

		sponsorID, err := sponsorsController.Add(mockSponsor)

		if err != nil {
			t.Error(err)
		}

		firstSponsor := testUtils.GetInterfaceAtCollectionRow(
			db,
			"sponsors",
			reflect.TypeOf(models.Sponsor{}),
			0,
		).(models.Sponsor)

		if firstSponsor.ID.IsZero() {
			t.Error("Expected sponsor id to return a non zero ID")
			t.SkipNow()
		}

		if firstSponsor.ID.Hex() != sponsorID {
			t.Error("Expected stored sponsor's ID to be '" + sponsorID + "' but got '" + firstSponsor.ID.Hex() + "'")
		}
	})
}
