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
	Register(c *fiber.Ctx) error
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

	token := controller.AuthService.Login(c.Context(), request)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  services.ExpiresTime,
	})

	response := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "LOGIN SUCCESSFUL",
		Data:   nil,
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
		Status: "LOGOUT SUCCESSFUL",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (controller *AuthControllerImpl) Register(c *fiber.Ctx) error {
	request := web.RegisterRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	user := controller.AuthService.RegisterUser(c.Context(), request)

	response := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "REGISTER SUCCESSFUL",
		Data:   user,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
