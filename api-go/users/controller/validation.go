package controller

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func getValidator() *validator.Validate {
	var validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.RegisterValidation("type-string", isString)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("isUUID", isValidUUID)
	if err != nil {
		panic(err)
	}

	return validate
}

func isValidUUID(fl validator.FieldLevel) bool {
	err := uuid.Validate(fl.Field().String())
	isValid := err == nil
	return isValid
}

func isString(fl validator.FieldLevel) bool {
	field := fl.Field()
	// Check if the field's kind is reflect.String
	return field.Kind() == reflect.String
}

// instead of ValidateMap to error on unknown keys
func validateMapCustom(validate *validator.Validate, body map[string]any, rulesMap rules) error {
	for key, value := range body {
		rule, ok := rulesMap[key]
		if !ok {
			return errors.New("Invalid key: " + key)
		}
		err := validate.Var(value, rule)
		if err != nil {
			return err
		}

	}
	return nil
}
