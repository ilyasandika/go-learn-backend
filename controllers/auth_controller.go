package controllers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"uaspw2/helper"
	"uaspw2/models/web/request"
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

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "login successfully", nil)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *AuthControllerImpl) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "logout successfully", nil)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *AuthControllerImpl) Register(c *fiber.Ctx) error {
	req := request.RegisterRequest{}
	err := c.BodyParser(&req)
	helper.PanicIfErr(err)

	user := controller.AuthService.RegisterUser(c.Context(), req)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "register successfully", user)
	return c.Status(webResponse.Code).JSON(webResponse)
}
