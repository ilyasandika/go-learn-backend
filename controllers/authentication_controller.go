package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	web2 "uaspw2/models/web"
	"uaspw2/services"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	services.AuthServices
}

func NewAuthenticationController(authServices services.AuthServices) AuthController {
	return &AuthControllerImpl{
		AuthServices: authServices,
	}
}

func (controller *AuthControllerImpl) Login(c *fiber.Ctx) error {
	request := web2.LoginRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	token, err := controller.AuthServices.Login(c.Context(), request)
	helper.PanicIfErr(err)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Expires:  services.ExpiresTime,
	})

	response := web2.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "Login Successful",
		Data:   token,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
