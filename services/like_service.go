package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"uaspw2/helper"
	"uaspw2/models/entity"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/repositories"
)

type LikeService interface {
	Create(ctx context.Context, request request.LikeRequest) response.LikeResponse
	Delete(ctx context.Context, request request.LikeRequest)
	FindByArticleID(ctx context.Context, articleId int) []response.LikeResponse
	FindByUserID(ctx context.Context, userId int) []response.LikeResponse
}

type LikeServiceImpl struct {
	repositories.LikeRepository
	*sql.DB
	*validator.Validate
}

func NewLikeService(likeRepository repositories.LikeRepository, db *sql.DB, validate *validator.Validate) LikeService {
	return &LikeServiceImpl{
		LikeRepository: likeRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *LikeServiceImpl) Create(ctx context.Context, request request.LikeRequest) response.LikeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	req := entity.Like{
		UserId:    request.UserId,
		ArticleId: request.ArticleId,
	}

	like := service.LikeRepository.Create(ctx, tx, req)
	log.Info(like)
	like, err = service.LikeRepository.FindByArticleAndUser(ctx, tx, request.ArticleId, request.UserId)
	helper.PanicIfNotFound(err, "like not found")

	return helper.ToLikeResponse(like)
}

func (service *LikeServiceImpl) Delete(ctx context.Context, request request.LikeRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.LikeRepository.FindByArticleAndUser(ctx, tx, request.ArticleId, request.UserId)
	helper.PanicIfNotFound(err, "like not found")

	service.LikeRepository.Delete(ctx, tx, request.ArticleId, request.UserId)
}

func (service *LikeServiceImpl) FindByArticleID(ctx context.Context, articleId int) []response.LikeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	likes := service.LikeRepository.FindByArticleId(ctx, tx, articleId)
	return helper.ToLikeResponses(likes)
}

func (service *LikeServiceImpl) FindByUserID(ctx context.Context, userId int) []response.LikeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	likes := service.LikeRepository.FindByUserId(ctx, tx, userId)
	return helper.ToLikeResponses(likes)
}
