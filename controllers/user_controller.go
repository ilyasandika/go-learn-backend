package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/services"
)

type UserController interface {
	UpdateByToken(c *fiber.Ctx) error
	UpdateByPath(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindByPath(c *fiber.Ctx) error
	FindByToken(c *fiber.Ctx) error
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

func (controller *UserControllerImpl) UpdateByPath(c *fiber.Ctx) error {
	req := request.UserUpdateRequest{}
	err := c.BodyParser(&req)
	helper.PanicIfErr(err)

	req.Id = helper.ToIntFromParams(c.Params("userId"))

	data := controller.service.Update(c.Context(), req)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "user updated successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserControllerImpl) UpdateByToken(c *fiber.Ctx) error {
	req := request.UserUpdateRequest{}
	err := c.BodyParser(&req)
	helper.PanicIfErr(err)

	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	if req.Role != user.Role && user.Role != "admin" {
		panic(exception.NewInvalidCredentialsError("only admin can update role"))
	}

	req.Id = user.Id

	data := controller.service.Update(c.Context(), req)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "user updated successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserControllerImpl) Delete(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	controller.service.Delete(c.Context(), userId)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "user deleted successfully", nil)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserControllerImpl) FindByPath(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	data := controller.service.FindByID(c.Context(), userId)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "user found", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserControllerImpl) FindByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	userId := user.Id

	data := controller.service.FindByID(c.Context(), userId)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "user found", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	data := controller.service.FindAll(c.Context())

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "user list retrieved successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}
