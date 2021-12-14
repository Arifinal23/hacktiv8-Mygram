package validation

import (
	"github.com/arfan21/golang-mygram/model/modelcomment"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateComment(data modelcomment.Request) error {
	return validation.Errors{
		"message":  validation.Validate(data.Message, validation.Required),
		"photo_id": validation.Validate(data.PhotoID, validation.Required),
	}.Filter()
}

func ValidateCommentUpdate(data modelcomment.RequestUpdate) error {
	return validation.Errors{
		"message": validation.Validate(data.Message, validation.Required),
	}.Filter()
}
