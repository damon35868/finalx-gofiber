package common

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func JWTErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusUnauthorized
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return HttpException(ctx, code, err.Error())
}
