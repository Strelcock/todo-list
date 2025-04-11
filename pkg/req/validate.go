package req

import "github.com/go-playground/validator/v10"

func validate[T any](payload T) error {
	valid := validator.New()
	err := valid.Struct(payload)
	return err
}
