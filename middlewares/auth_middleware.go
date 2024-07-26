package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"uaspw2/config"
	"uaspw2/helper"
	"uaspw2/models/web"
)

func AuthRequired(c *fiber.Ctx) error {
	_, err := helper.VerifyToken(c, jwt.MapClaims{})

	if err != nil {
		return helper.HandleTokenError(c)
	}

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	userClaims := &config.UserClaims{}
	token, err := helper.VerifyToken(c, userClaims)

	if err != nil {
		return helper.HandleTokenError(c)
	}

	if claims, ok := token.(*config.UserClaims); ok && claims.Role == "admin" {
		return c.Next()
	}

	return helper.HandleTokenError(c)
}

func UserOnly(c *fiber.Ctx) error {
	userClaims := &config.UserClaims{}
	token, err := helper.VerifyToken(c, userClaims)

	if err != nil {
		return helper.HandleTokenError(c)
	}

	if claims, ok := token.(*config.UserClaims); ok && claims.Role == "user" {
		return c.Next()
	}
	return helper.HandleTokenError(c)
}

func GuestOnly(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString != "" {
		errorResponse := web.ErrorResponse{
			Code:   fiber.StatusForbidden,
			Status: "FORBIDDEN",
		}
		return c.Status(fiber.StatusForbidden).JSON(errorResponse)
	}
	return c.Next()
}
