package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	web2 "uaspw2/models/web"
	"uaspw2/services"
)

type AuthenticationController interface {
	Login(c *fiber.Ctx) error
}

type AuthenticationControllerImpl struct {
	services.AuthenticationServices
}

func NewAuthenticationController(authenticationServices services.AuthenticationServices) AuthenticationController {
	return &AuthenticationControllerImpl{
		AuthenticationServices: authenticationServices,
	}
}

func (controller *AuthenticationControllerImpl) Login(c *fiber.Ctx) error {
	request := web2.LoginRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	token, err := controller.AuthenticationServices.Login(c.Context(), request)
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
