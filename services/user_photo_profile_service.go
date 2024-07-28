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

type UserProfilePhotoService interface {
	UpdateByUserId(ctx context.Context, request request.UserProfilePhotoRequest) response.UserProfilePhotoResponse
	FindByUserId(ctx context.Context, userId int) response.UserProfilePhotoResponse
}

type UserProfilePhotoServiceImpl struct {
	UserProfilePhotoRepository repositories.UserProfilePhotoRepository
	DB                         *sql.DB
	Validate                   *validator.Validate
}

func NewUserProfilePhotoService(userProfilePhotoRepository repositories.UserProfilePhotoRepository, DB *sql.DB, validate *validator.Validate) UserProfilePhotoService {
	return &UserProfilePhotoServiceImpl{
		UserProfilePhotoRepository: userProfilePhotoRepository,
		DB:                         DB,
		Validate:                   validate,
	}
}

func (service *UserProfilePhotoServiceImpl) UpdateByUserId(ctx context.Context, request request.UserProfilePhotoRequest) response.UserProfilePhotoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	userProfilePhoto := entity.UserProfilePhoto{
		UserId: request.UserId,
		Path:   request.Path,
	}

	result := service.UserProfilePhotoRepository.Update(ctx, tx, userProfilePhoto)
	return helper.ToUserPhotoProfileResponse(result)
}

func (service *UserProfilePhotoServiceImpl) FindByUserId(ctx context.Context, userId int) response.UserProfilePhotoResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.UserProfilePhotoRepository.FindByUserID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToUserPhotoProfileResponse(result)
}
