package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/services"
)

type UserProfileController interface {
	FindAll(c *fiber.Ctx) error
	FindByToken(c *fiber.Ctx) error
	FindByPath(c *fiber.Ctx) error
	UpdateByToken(c *fiber.Ctx) error
}

type UserProfileControllerImpl struct {
	UserProfileService services.UserProfileService
}

func NewUserProfileController(userProfileService services.UserProfileService) UserProfileController {
	return &UserProfileControllerImpl{
		UserProfileService: userProfileService,
	}
}

func (controller *UserProfileControllerImpl) FindAll(c *fiber.Ctx) error {
	data := controller.UserProfileService.FindAll(c.Context())

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "User profile list retrieved successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserProfileControllerImpl) FindByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	data := controller.UserProfileService.FindByUserID(c.Context(), user.Id)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "User profile found", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserProfileControllerImpl) FindByPath(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	data := controller.UserProfileService.FindByUserID(c.Context(), userId)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "User profile found", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *UserProfileControllerImpl) UpdateByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	req := request.UserProfileUpdateRequest{}
	err = c.BodyParser(&req)
	helper.PanicIfErr(err)

	req.UserId = user.Id

	data := controller.UserProfileService.Update(c.Context(), req)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "User profile updated successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}
