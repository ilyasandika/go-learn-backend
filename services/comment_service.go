package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/models/entity"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/repositories"
)

type CommentService interface {
	Create(ctx context.Context, request request.CommentRequest) response.CommentResponse
	FindByArticleID(ctx context.Context, articleId int) []response.CommentResponse
	Delete(ctx context.Context, commentId int, userId int)
}

type CommentServiceImpl struct {
	repositories.CommentRepository
	*sql.DB
	*validator.Validate
}

func NewCommentService(commentRepository repositories.CommentRepository, db *sql.DB, validate *validator.Validate) *CommentServiceImpl {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (controller *CommentServiceImpl) Create(ctx context.Context, request request.CommentRequest) response.CommentResponse {
	err := controller.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := controller.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	req := entity.Comment{
		UserId:    request.UserId,
		ArticleId: request.ArticleId,
		Comment:   request.Comment,
	}

	comment := controller.CommentRepository.Create(ctx, tx, req)
	comment, err = controller.CommentRepository.FindByID(ctx, tx, comment.Id)
	helper.PanicIfNotFound(err, "comment not found")

	return helper.ToCommentResponse(comment)

}

func (controller *CommentServiceImpl) FindByArticleID(ctx context.Context, articleId int) []response.CommentResponse {
	tx, err := controller.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	comments := controller.CommentRepository.FindByArticleID(ctx, tx, articleId)
	log.Info(comments)

	return helper.ToCommentResponses(comments)
}

func (controller *CommentServiceImpl) Delete(ctx context.Context, commentId int, userId int) {
	tx, err := controller.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	comment, err := controller.CommentRepository.FindByID(ctx, tx, commentId)
	helper.PanicIfNotFound(err, "comment not found")

	if comment.UserId != userId {
		panic(exception.NewInvalidCredentialsError("you cannot delete another user comment"))
	}

	controller.CommentRepository.Delete(ctx, tx, comment.Id, userId)

}
