package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/services"
)

type LikeController interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindByArticleId(c *fiber.Ctx) error
	FindByUserId(c *fiber.Ctx) error
}

type likeControllerImpl struct {
	services.LikeService
}

func NewLikeController(likeService services.LikeService) LikeController {
	return &likeControllerImpl{
		LikeService: likeService,
	}
}

func (controller *likeControllerImpl) Create(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	articleId := helper.ToIntFromParams(c.Params("articleId"))

	req := request.LikeRequest{
		ArticleId: articleId,
		UserId:    user.Id,
	}
	data := controller.LikeService.Create(c.Context(), req)

	webResponse := helper.CreateSuccessResponse(fiber.StatusCreated, "like successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *likeControllerImpl) Delete(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	articleId := helper.ToIntFromParams(c.Params("articleId"))

	req := request.LikeRequest{
		ArticleId: articleId,
		UserId:    user.Id,
	}
	controller.LikeService.Delete(c.Context(), req)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "unlike successfully", nil)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *likeControllerImpl) FindByArticleId(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))

	data := controller.LikeService.FindByArticleID(c.Context(), articleId)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "like list by article retrieved successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *likeControllerImpl) FindByUserId(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	data := controller.LikeService.FindByUserID(c.Context(), userId)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "like list by user retrieved successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}
