package sponsors

import (
	"regexp"

	"github.com/ionian-uni-ieee/ieeesb-app/internal/app/models"
	"github.com/ionian-uni-ieee/ieeesb-app/pkg/validation"
)

// Validates if a sponsor object is valid
func (s *Service) Validate(sponsor models.Sponsor) *validation.Validation {
	v := &validation.Validation{}

	if sponsor.GetName() == "" {
		v.AddError("name", ErrNameEmpty)
	}

	emailsRegex := regexp.MustCompile("^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$")
	emails := sponsor.GetEmails()
	for email := range emails {
		isEmailValid := emailsRegex.MatchString(emails[email])
		if !isEmailValid {
			v.AddError("emails", ErrInvalidEmails)
			break
		}
	}

	phonesRegex := regexp.MustCompile("^[0-9]{10}$")
	phones := sponsor.GetPhones()
	for phone := range phones {
		isPhoneValid := phonesRegex.MatchString(phones[phone])
		if !isPhoneValid {
			v.AddError("phones", ErrInvalidPhones)
			break
		}
	}
	
	return v
}