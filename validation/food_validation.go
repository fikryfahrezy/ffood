package validation

import (
	"encoding/json"

	"github.com/fikryfahrezy/ffood/exception"
	"github.com/fikryfahrezy/ffood/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func InsertFoodValidation(Request model.InsertFoodRequest) (err error) {
	err = validation.ValidateStruct(&Request,
		validation.Field(&Request.Name, validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		return exception.ValidationError{
			Message: string(b),
		}
	}

	return nil
}
