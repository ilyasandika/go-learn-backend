package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/services"
)

type ArticleController interface {
	CreateByToken(c *fiber.Ctx) error
	UpdateByID(c *fiber.Ctx) error
	DeleteByID(c *fiber.Ctx) error
	FindByUserId(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindAllPublished(c *fiber.Ctx) error
	FindAllPublishedByUserID(c *fiber.Ctx) error
	FindPublishedByID(c *fiber.Ctx) error
	FindAllUnpublished(c *fiber.Ctx) error
	FindAllUnPublishedByUserID(c *fiber.Ctx) error
}

type ArticleControllerImpl struct {
	services.ArticleService
}

func NewArticleController(articleService services.ArticleService) ArticleController {
	return &ArticleControllerImpl{
		ArticleService: articleService,
	}
}

func (controller *ArticleControllerImpl) CreateByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	req := request.ArticleCreateRequest{}
	err = c.BodyParser(&req)
	helper.PanicIfErr(err)

	req.UserId = user.Id
	req.IsPublished = false

	article := controller.ArticleService.Create(c.Context(), req)

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusCreated,
		Message: "article created successfully",
		Data:    article,
	}
	return c.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ArticleControllerImpl) UpdateByID(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	articleId := helper.ToIntFromParams(c.Params("articleId"))

	req := request.ArticleUpdateRequest{}
	err = c.BodyParser(&req)
	helper.PanicIfErr(err)

	req.UserId = user.Id
	req.Id = articleId

	if user.Role != "admin" && req.IsPublished == true {
		webResponse := response.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
			Error:   "You are not allowed to publish this article",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(webResponse)
	}

	log.Info(req)

	article := controller.ArticleService.Update(c.Context(), req)

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "article updated successfully",
		Data:    article,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) DeleteByID(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))

	controller.ArticleService.Delete(c.Context(), articleId)

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "article deleted successfully",
		Data:    nil,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindByUserId(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	articles := controller.ArticleService.FindByUserID(c.Context(), userId)
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "articles list by user retrieved successfully",
		Data:    articles,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindByID(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))

	articles := controller.ArticleService.FindByID(c.Context(), articleId)
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "article found",
		Data:    articles,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAll(c *fiber.Ctx) error {
	articles := controller.ArticleService.FindAll(c.Context())
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "articles list retrieved successfully",
		Data:    articles,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllPublished(c *fiber.Ctx) error {
	articles := controller.ArticleService.FindAllPublished(c.Context())
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "published articles list retrieved successfully",
		Data:    articles,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllPublishedByUserID(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	articles := controller.ArticleService.FindAllPublishedByUserID(c.Context(), userId)
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "published articles list by user retrieved successfully",
		Data:    articles,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindPublishedByID(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))
	article := controller.ArticleService.FindPublishedByID(c.Context(), articleId)
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "published article found",
		Data:    article,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllUnpublished(c *fiber.Ctx) error {
	articles := controller.ArticleService.FindAllUnpublished(c.Context())
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "unpublished articles list retrieved successfully",
		Data:    articles,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllUnPublishedByUserID(c *fiber.Ctx) error {
	userId := helper.ToIntFromParams(c.Params("userId"))

	articles := controller.ArticleService.FindAllUnpublishedByUserID(c.Context(), userId)
	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "published articles list by user retrieved successfully",
		Data:    articles,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}
