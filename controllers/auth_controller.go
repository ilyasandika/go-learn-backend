package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
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
	req := request.LoginRequest{}
	err := c.BodyParser(&req)
	helper.PanicIfErr(err)

	token := controller.AuthService.Login(c.Context(), req)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  services.TokenExpiresTime,
	})

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Login successfully",
		Data:    nil,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *AuthControllerImpl) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Logout successfully",
		Data:    nil,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *AuthControllerImpl) Register(c *fiber.Ctx) error {
	req := request.RegisterRequest{}
	err := c.BodyParser(&req)
	helper.PanicIfErr(err)

	user := controller.AuthService.RegisterUser(c.Context(), req)

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Register successfully",
		Data:    user,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}
