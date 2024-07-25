package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	web2 "uaspw2/models/web"
	"uaspw2/services"
)

type UserController interface {
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	request := web2.UserUpdateRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	request.Id = helper.ToIntFromParams(c.Params("userId"))

	data := controller.service.Update(c.Context(), request)
	successResponse := web2.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	controller.service.Delete(c.Context(), userId)

	successResponse := web2.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) FindById(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	data := controller.service.FindByID(c.Context(), userId)

	successResponse := web2.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	data := controller.service.FindAll(c.Context())

	successResponse := web2.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}
