package validation

import (
	"github.com/arfan21/golang-mygram/model/modelsocialmedia"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateSocialMediaCreate(data modelsocialmedia.Request) error {
	return validation.Errors{
		"name":             validation.Validate(data.Name, validation.Required),
		"social_media_url": validation.Validate(data.SocialMediaUrl, validation.Required),
	}.Filter()
}
