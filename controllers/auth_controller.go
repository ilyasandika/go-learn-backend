package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"uaspw2/helper"
	"uaspw2/models/web"
	"uaspw2/services"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	services.AuthService
}

func NewAuthenticationController(authServices services.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authServices,
	}
}

func (controller *AuthControllerImpl) Login(c *fiber.Ctx) error {
	request := web.LoginRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	token, err := controller.AuthService.Login(c.Context(), request)
	helper.PanicIfErr(err)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  services.ExpiresTime,
	})

	response := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "Login Successful",
		Data:   token,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *AuthControllerImpl) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	response := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "Logout Successful",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
