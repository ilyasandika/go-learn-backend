package controller

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	"uaspw2/services"
	"uaspw2/web"
)

type UserController interface {
	Create(c *fiber.Ctx) error
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

func (controller *UserControllerImpl) Create(c *fiber.Ctx) error {
	request := web.UserCreateRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	data := controller.service.Create(c.Context(), request)
	successResponse := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	request := web.UserUpdateRequest{}
	err := c.BodyParser(&request)
	helper.PanicIfErr(err)

	request.Id = helper.ToIntFromParams(c.Params("userId"))

	data := controller.service.Update(c.Context(), request)
	successResponse := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	controller.service.Delete(c.Context(), userId)

	successResponse := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   nil,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) FindById(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	data := controller.service.FindByID(c.Context(), userId)

	successResponse := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}

func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	data := controller.service.FindAll(c.Context())

	successResponse := web.SuccessResponse{
		Code:   fiber.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}

	return c.Status(fiber.StatusOK).JSON(successResponse)
}
