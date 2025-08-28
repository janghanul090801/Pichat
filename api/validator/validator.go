package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	_ = Validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		s := fl.Field().String()

		hasLetter := regexp.MustCompile("[A-Za-z]").MatchString(s)
		hasDigit := regexp.MustCompile("[0-9]").MatchString(s)
		hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(s)
		longEnough := len(s) >= 8

		return hasLetter && hasDigit && hasSpecial && longEnough
	})
}
