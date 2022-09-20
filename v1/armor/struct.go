package armor

import "github.com/go-playground/validator/v10"

func ValidateStruct(model interface{}) error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		return err
	}

	return nil
}
