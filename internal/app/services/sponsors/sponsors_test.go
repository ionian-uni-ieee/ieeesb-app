package sponsors_test

import (
	"testing"
	
	testingDb "github.com/ionian-uni-ieee/ieeesb-app/internal/app/drivers/database/testing"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/repositories"
	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/services/sponsors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var database = testingDb.MakeDatabaseDriver()
var reps = repositories.MakeRepositories(database)
var service = sponsors.MakeService(reps)

func TestValidate(t *testing.T) {
	t.Run("Should return true for valid sponsor", func(t *testing.T){
		validSponsor := models.Sponsor{
			ID:				primitive.NewObjectID(),
			Name: 		"test",
			Emails: 	[]string{"test@test.com", "test@test2.com"},
			Phones: 	[]string{"6912345678", "2661012345"},
			Logo: 		"test",
		}
		sponsorIsValid := !service.Validate(validSponsor).HasError()
		if !sponsorIsValid {
			t.Error("Sponsor should be valid")
		}
	})
	t.Run("Should return false for wrong email form", func(t *testing.T){
		invalidSponsor := models.Sponsor{
			ID:				primitive.NewObjectID(),
			Name: 		"test",
			Emails: 	[]string{"wrongemailform", "test@test2.com"},
			Phones: 	[]string{"6912345678", "2661012345"},
			Logo: 		"test",
		}
		validation := *service.Validate(invalidSponsor)
		sponsorIsValid := !validation.HasError()
		if sponsorIsValid {
			t.Error("Sponsor should be invalid")
		}
		if validation["emails"] != sponsors.ErrInvalidEmails {
			t.Error("Expected \"" + sponsors.ErrInvalidEmails.Error() + "\" but got \"" + validation["email"].Error() + "\"")
		}
	})
	t.Run("Should return false for empty name", func(t *testing.T){
		invalidSponsor := models.Sponsor{
			ID:				primitive.NewObjectID(),
			Name: 		"",
			Emails: 	[]string{"test@test.com", "test@test2.com"},
			Phones: 	[]string{"6912345678", "2661012345"},
			Logo: 		"test",
		}
		validation := *service.Validate(invalidSponsor)
		sponsorIsValid := !validation.HasError()
		if sponsorIsValid {
			t.Error("Sponsor should be invalid")
		}
		if validation["name"] != sponsors.ErrNameEmpty {
			t.Error("Expected \"" + sponsors.ErrNameEmpty.Error() + "\" but got \"" + validation["name"].Error() + "\"")
		}
	})
	t.Run("Should return false for wrong phone form", func(t *testing.T){
		invalidSponsor := models.Sponsor{
			ID:				primitive.NewObjectID(),
			Name: 		"test",
			Emails: 	[]string{"test@test.com", "test@test2.com"},
			Phones: 	[]string{"test123", "2661012345"},
			Logo: 		"test",
		}
		validation := *service.Validate(invalidSponsor)
		sponsorIsValid := !validation.HasError()
		if sponsorIsValid {
			t.Error("Sponsor should be invalid")
		}
		if validation["phones"] != sponsors.ErrInvalidPhones {
			t.Error("Expected \"" + sponsors.ErrInvalidPhones.Error() + "\" but got \"" + validation["phones"].Error() + "\"")
		}
	})
}
