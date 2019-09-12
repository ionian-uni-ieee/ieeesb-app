package sponsors_test

import (
	"testing"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/testUtils"
)

func TestGetSponsors(t *testing.T) {
	// Setup
	db, sponsorsController := makeController()

	t.Run("Should return an array of 2 sponsors", func(t *testing.T) {
		testUtils.SetupData(db, "sponsors", mockSponsor, mockSponsor2)

		gotSponsors, err := sponsorsController.GetSponsors(0, 2)

		if err != nil {
			t.Error(err)
			t.SkipNow()
		}

		if len(gotSponsors) != 2 {
			t.Error("Expected sponsors array with length of 2, but got", len(gotSponsors))
			t.SkipNow()
		}

		gotValidData := gotSponsors[0].ID == mockSponsor.ID ||
			gotSponsors[1].ID == mockSponsor2.ID
		if !gotValidData {
			t.Error("Expected returned data to have the same IDs with stored data")
		}
	})
}
