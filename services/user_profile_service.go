package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/repositories"
)

type UserProfileService interface {
	Update(ctx context.Context, request request.UserProfileUpdateRequest) response.UserProfileResponse
	FindByUserID(ctx context.Context, userId int) response.UserProfileResponse
	FindAll(ctx context.Context) []response.UserProfileResponse
}

type UserProfileServiceImpl struct {
	UserProfileRepository repositories.UserProfileRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewUserProfileService(userProfileRepository repositories.UserProfileRepository, db *sql.DB, validate *validator.Validate) UserProfileService {
	return &UserProfileServiceImpl{
		UserProfileRepository: userProfileRepository,
		DB:                    db,
		Validate:              validate,
	}
}

func (service *UserProfileServiceImpl) Update(ctx context.Context, request request.UserProfileUpdateRequest) response.UserProfileResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	userProfile, err := service.UserProfileRepository.FindByUserID(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	data := service.UserProfileRepository.Update(ctx, tx, userProfile)

	return helper.ToUserProfileResponse(data)

}

func (service *UserProfileServiceImpl) FindByUserID(ctx context.Context, userId int) response.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	userProfile, err := service.UserProfileRepository.FindByUserID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserProfileResponse(userProfile)

}

func (service *UserProfileServiceImpl) FindAll(ctx context.Context) []response.UserProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	userProfile := service.UserProfileRepository.FindAll(ctx, tx)

	return helper.ToUserProfileResponses(userProfile)
}
