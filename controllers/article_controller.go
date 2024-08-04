package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/services"
)

type ArticleController interface {
	CreateByToken(c *fiber.Ctx) error
	UpdateByID(c *fiber.Ctx) error
	DeleteByID(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	FindAllPublished(c *fiber.Ctx) error
	FindAllPublishedByUserID(c *fiber.Ctx) error
	FindAllUnpublished(c *fiber.Ctx) error
	FindAllUnpublishedByUserID(c *fiber.Ctx) error
	PublishArticle(c *fiber.Ctx) error
	UnpublishArticle(c *fiber.Ctx) error
}

type ArticleControllerImpl struct {
	services.ArticleService
}

func NewArticleController(articleService services.ArticleService) ArticleController {
	return &ArticleControllerImpl{
		ArticleService: articleService,
	}
}

func (controller *ArticleControllerImpl) PublishArticle(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))
	controller.ArticleService.UpdateStatusArticleByID(c.Context(), articleId, true)
	return c.SendStatus(fiber.StatusOK)
}

func (controller *ArticleControllerImpl) UnpublishArticle(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))
	controller.ArticleService.UpdateStatusArticleByID(c.Context(), articleId, false)
	return c.SendStatus(fiber.StatusOK)
}

func (controller *ArticleControllerImpl) CreateByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	req := request.ArticleCreateRequest{}
	err = c.BodyParser(&req)
	helper.PanicIfErr(err)

	req.UserId = user.Id
	req.IsPublished = false

	mediaRequests := req.Media

	article := controller.ArticleService.Create(c.Context(), req, mediaRequests)

	webResponse := helper.CreateSuccessResponse(fiber.StatusCreated, "article created successfully", article)
	return c.Status(webResponse.Code).JSON(webResponse)
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
		webResponse := helper.CreateErrorResponse(fiber.StatusUnauthorized, "unauthorized", "you are not allowed to publish this article")
		return c.Status(webResponse.Code).JSON(webResponse)
	}

	log.Info(req)

	article := controller.ArticleService.Update(c.Context(), req)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "article updated successfully", article)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *ArticleControllerImpl) DeleteByID(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))

	controller.ArticleService.Delete(c.Context(), articleId)

	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "article deleted successfully", nil)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindByID(c *fiber.Ctx) error {
	articleId := helper.ToIntFromParams(c.Params("articleId"))

	articles := controller.ArticleService.FindByID(c.Context(), articleId)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "article found", articles)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllPublished(c *fiber.Ctx) error {
	articles := controller.ArticleService.FindAllPublished(c.Context())
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "published articles list retrieved successfully", articles)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllPublishedByUserID(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	articles := controller.ArticleService.FindAllPublishedByUserID(c.Context(), user.Id)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "published articles list by user retrieved successfully", articles)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllUnpublished(c *fiber.Ctx) error {
	articles := controller.ArticleService.FindAllUnpublished(c.Context())
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "unpublished articles list retrieved successfully", articles)
	return c.Status(webResponse.Code).JSON(webResponse)
}

func (controller *ArticleControllerImpl) FindAllUnpublishedByUserID(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	articles := controller.ArticleService.FindAllUnpublishedByUserID(c.Context(), user.Id)
	webResponse := helper.CreateSuccessResponse(fiber.StatusOK, "unpublished articles list by user retrieved successfully", articles)
	return c.Status(webResponse.Code).JSON(webResponse)
}
