package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/models/entity"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/repositories"
)

type ArticleService interface {
	Create(ctx context.Context, request request.ArticleCreateRequest, mediaRequests []request.ArticleMediaCreateRequest) response.ArticleResponse
	Update(ctx context.Context, request request.ArticleUpdateRequest) response.ArticleResponse
	Delete(ctx context.Context, articleId int)
	FindByID(ctx context.Context, articleId int) response.ArticleResponse
	FindAllPublished(ctx context.Context) []response.ArticleResponse
	FindAllPublishedByUserID(ctx context.Context, userId int) []response.ArticleResponse
	FindAllUnpublished(ctx context.Context) []response.ArticleResponse
	FindAllUnpublishedByUserID(ctx context.Context, userId int) []response.ArticleResponse
	UpdateStatusArticleByID(ctx context.Context, articleId int, status bool)
}

type ArticleServiceImpl struct {
	repositories.ArticleRepository
	*sql.DB
	*validator.Validate
}

func NewArticleService(articleRepository repositories.ArticleRepository, db *sql.DB, validate *validator.Validate) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepository: articleRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ArticleServiceImpl) UpdateStatusArticleByID(ctx context.Context, articleId int, status bool) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)
	service.ArticleRepository.UpdatePublishStatus(ctx, tx, articleId, status)
}

func (service *ArticleServiceImpl) Create(ctx context.Context, request request.ArticleCreateRequest, mediaRequests []request.ArticleMediaCreateRequest) response.ArticleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	req := entity.Article{
		UserId:      request.UserId,
		Title:       request.Title,
		Description: request.Description,
		Content:     request.Content,
		IsPublished: request.IsPublished,
	}

	data := service.ArticleRepository.Create(ctx, tx, req)
	article, err := service.ArticleRepository.FindByID(ctx, tx, data.Id)

	helper.PanicIfErr(err)

	for _, media := range mediaRequests {
		mediaReq := entity.ArticleMedia{
			ArticleId: data.Id,
			Type:      media.Type,
			Path:      media.Path,
		}
		service.ArticleRepository.CreateMedia(ctx, tx, mediaReq)
	}

	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) Update(ctx context.Context, request request.ArticleUpdateRequest) response.ArticleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	req := entity.Article{
		Id:          request.Id,
		UserId:      request.UserId,
		Title:       request.Title,
		Description: request.Description,
		IsPublished: false,
	}

	article, err := service.ArticleRepository.FindByID(ctx, tx, req.Id)
	helper.PanicIfNotFound(err, "article not found")

	if article.UserId != req.UserId {
		panic(exception.NewInvalidCredentialsError("unauthorized"))
	}

	articleResponse := service.ArticleRepository.Update(ctx, tx, req)

	articleResponse.CreatedAt = article.CreatedAt
	articleResponse.UpdatedAt = article.UpdatedAt

	return helper.ToArticleResponse(articleResponse)
}

func (service *ArticleServiceImpl) Delete(ctx context.Context, articleId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindByID(ctx, tx, articleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.ArticleRepository.Delete(ctx, tx, article.Id)
}

func (service *ArticleServiceImpl) FindByID(ctx context.Context, articleId int) response.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	article, err := service.ArticleRepository.FindByID(ctx, tx, articleId)
	helper.PanicIfNotFound(err, "article not found")

	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) FindAllPublished(ctx context.Context) []response.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepository.FindAllByPublishStatus(ctx, tx, true)

	return helper.ToArticleResponses(articles)
}

func (service *ArticleServiceImpl) FindAllPublishedByUserID(ctx context.Context, userId int) []response.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepository.FindAllByPublishStatusAndUserID(ctx, tx, true, userId)

	return helper.ToArticleResponses(articles)
}

func (service *ArticleServiceImpl) FindAllUnpublished(ctx context.Context) []response.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepository.FindAllByPublishStatus(ctx, tx, false)

	return helper.ToArticleResponses(articles)
}

func (service *ArticleServiceImpl) FindAllUnpublishedByUserID(ctx context.Context, userId int) []response.ArticleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	articles := service.ArticleRepository.FindAllByPublishStatusAndUserID(ctx, tx, false, userId)

	return helper.ToArticleResponses(articles)
}
