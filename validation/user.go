package validation

import (
	"errors"

	"github.com/arfan21/golang-mygram/model/modeluser"
	"github.com/arfan21/golang-mygram/repository/repositoryuser"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func isEmailExist(repo repositoryuser.RepositoryUser) validation.RuleFunc {
	return func(value interface{}) error {
		email, ok := value.(string)
		if !ok {
			return errors.New("invalid email address")
		}

		return repo.IsEmailExist(email)
	}
}

func ValidateUserCreate(data modeluser.Request, repo repositoryuser.RepositoryUser) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email, validation.By(isEmailExist(repo))),
		"username": validation.Validate(data.Username, validation.Required),
		"password": validation.Validate(data.Password, validation.Required, validation.Length(8, 20)),
		"age":      validation.Validate(data.Age, validation.Required),
	}.Filter()
}

func ValidateUserLogin(data modeluser.RequestLogin) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email),
		"password": validation.Validate(data.Password, validation.Required, validation.Length(8, 20).Error("invalid email or password")),
	}.Filter()
}

func ValidateUserUpdate(data modeluser.Request) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email),
		"username": validation.Validate(data.Username, validation.Required),
	}.Filter()
}
