package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"uaspw2/config"
	"uaspw2/models/web"
)

func VerifyToken(c *fiber.Ctx, claims jwt.Claims) (jwt.Claims, error) {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return nil, fiber.ErrUnauthorized
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims, nil
}

func HandleTokenError(c *fiber.Ctx) error {
	errorResponse := web.ErrorResponse{
		Code:   fiber.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
}

func GetUserByToken(c *fiber.Ctx) (web.UserUpdateRequest, error) {
	userClaims := &config.UserClaims{}
	token, err := VerifyToken(c, userClaims)
	user := web.UserUpdateRequest{}

	if err != nil {
		return user, HandleTokenError(c)
	}

	if claims, ok := token.(*config.UserClaims); ok {
		user.Id = claims.Id
		user.Username = claims.Username
		user.Role = claims.Role
		return user, nil
	}

	return user, HandleTokenError(c)
}
