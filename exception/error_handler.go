package exception

import (
	"encoding/json"
	"golang-simple-boilerplate/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		var obj interface{}
		_ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   struct{}{},
			Error:  obj,
		})
	}

	if err == gorm.ErrRecordNotFound {
		return ctx.Status(404).JSON(model.Response{
			Code:   404,
			Status: "NOT_FOUND",
			Data:   nil,
			Error:  nil,
		})
	}

	if err.Error() == "PHONE_REGISTERED" {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"phone_number": "MUST_UNIQUE",
			},
		})
	}

	if err.Error() == "NIK_REGISTERED" {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"nik": "MUST_UNIQUE",
			},
		})
	}

	if err.Error() == "USERNAME_REGISTERED" {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"username": "MUST_UNIQUE",
			},
		})
	}

	if err.Error() == "EMAIL_REGISTERED" {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"email": "MUST_UNIQUE",
			},
		})
	}

	if err.Error() == "PASSWORD_OLD_NOTMATCH" {
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"old_password": "PASSWORD_OLD_NOTMATCH",
			},
		})
	}

	return ctx.Status(500).JSON(model.Response{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
