package validation

import (
	"github.com/arfan21/golang-mygram/model/modelphoto"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidatePhotoCreate(data modelphoto.Request) error {
	return validation.Errors{
		"title":     validation.Validate(data.Title, validation.Required),
		"caption":   validation.Validate(data.Caption),
		"photo_url": validation.Validate(data.PhotoURL, validation.Required),
		"userId":    validation.Validate(data.UserID),
	}.Filter()
}
