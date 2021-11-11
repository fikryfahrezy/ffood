package validation

import (
	"encoding/json"
	"github.com/fikryfahrezy/ffood/exception"
	"github.com/fikryfahrezy/ffood/model"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func LoginValidation(Request model.AuthRequest) (err error) {
	err = validation.ValidateStruct(&Request,
		validation.Field(&Request.Password, validation.Required.Error("NOT_BLANK")),
		validation.Field(&Request.Email, validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		return exception.ValidationError{
			Message: string(b),
		}
	}

	return nil
}

func RegisterValidation(Request model.RegisterRequest) (err error) {
	err = validation.ValidateStruct(&Request,
		validation.Field(&Request.Email, validation.Required.Error("NOT_BLANK"), is.Email),
		validation.Field(&Request.Name, validation.Required.Error("NOT_BLANK")),
		validation.Field(&Request.Password, validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		return exception.ValidationError{
			Message: string(b),
		}
	}

	return nil
}
