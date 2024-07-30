package controllers

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/services"
)

type CommentController interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindByArticleId(c *fiber.Ctx) error
}

type CommentControllerImpl struct {
	services.CommentService
}

func NewCommentController(commentService services.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (controller *CommentControllerImpl) Create(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	articleId := helper.ToIntFromParams(c.Params("articleId"))

	var commentRequest request.CommentRequest
	err = c.BodyParser(&commentRequest)
	helper.PanicIfErr(err)

	commentRequest.UserId = user.Id
	commentRequest.ArticleId = articleId

	data := controller.CommentService.Create(c.Context(), commentRequest)

	webResponse := helper.CreateSuccessResponse(fiber.StatusCreated, "comment created successfully", data)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *CommentControllerImpl) Delete(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	commentId := helper.ToIntFromParams(c.Params("commentId"))

	controller.CommentService.Delete(c.Context(), commentId, user.Id)

	webResponse := helper.CreateSuccessResponse(fiber.StatusCreated, "comment deleted successfully", nil)

	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *CommentControllerImpl) FindByArticleId(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))
	comments := controller.CommentService.FindByArticleID(c.Context(), articleId)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "comment list by article retrieved successfully", comments)
	return c.Status(webResponse.Code).JSON(webResponse)
}
