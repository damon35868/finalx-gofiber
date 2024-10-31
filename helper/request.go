package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func ParamHandler[T any](c *fiber.Ctx, method string, req *T) error {
	if method == "GET" || method == "get" {
		if err := c.ParamsParser(&req); err != nil {
			return err
		}
	}
	if method == "POST" || method == "post" {
		if err := c.BodyParser(&req); err != nil {
			return err
		}
	}
	if err := Validate.Struct(req); err != nil {
		return err
	}
	return nil
}
