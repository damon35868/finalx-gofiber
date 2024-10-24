package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ResponseOptions struct {
	Code    int
	Message *string
}

func GetTokenUserId(c *fiber.Ctx) int32 {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	if userIdFloat, ok := userId.(float64); ok {
		return int32(userIdFloat)
	}
	if userIdStr, ok := userId.(string); ok {
		if userIdInt, err := strconv.Atoi(userIdStr); err == nil {
			return int32(userIdInt)
		}
	}
	return 0
}

func HttpException(c *fiber.Ctx, errCode int, msgs ...string) error {
	code, msg := fiber.StatusInternalServerError, "未知错误"
	if errCode != 0 {
		code = errCode
	}
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return c.Status(code).JSON(&fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func Response(ctx *fiber.Ctx, data interface{}, options ...ResponseOptions) error {
	code, message := fiber.StatusOK, "请求成功"

	if len(options) > 0 {
		if options[0].Code != 0 {
			code = options[0].Code
		}
		if options[0].Message != nil {
			message = *options[0].Message
		}
	}

	return ctx.JSON(&fiber.Map{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func HasNextPage(skip, take, total int) bool {
	return skip*take < total
}
