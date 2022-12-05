package utilx

import validator "github.com/go-playground/validator/v10"

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
