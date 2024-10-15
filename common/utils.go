package common

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func MD5(s string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x\n", sum)
}

func GetTokenUserId(c *fiber.Ctx) int {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	if userIdFloat, ok := userId.(float64); ok {
		return int(userIdFloat)
	}
	if userIdStr, ok := userId.(string); ok {
		if userIdInt, err := strconv.Atoi(userIdStr); err == nil {
			return userIdInt
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

func Response(ctx *fiber.Ctx, data interface{}, codes ...int) error {
	code := fiber.StatusOK
	if len(codes) > 0 {
		code = codes[0]
	}

	return ctx.JSON(&fiber.Map{
		"code":    code,
		"message": "请求成功",
		"data":    data,
	})
}
