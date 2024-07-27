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

type UserService interface {
	Update(ctx context.Context, request request.UserUpdateRequest) response.UserResponse
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewUserService(userRepository repositories.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		validate:       validate,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) response.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	var user entity.User
	user, err = service.UserRepository.FindByID(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Username = request.Username
	user.Role = request.Role

	if request.Password != "" {
		hashedPassword, err := helper.HashPassword(request.Password)
		helper.PanicIfErr(err)
		user.Password = hashedPassword
	}

	user = service.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	var user entity.User
	user, err = service.UserRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user.Id)
}

func (service *UserServiceImpl) FindByID(ctx context.Context, id int) response.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	var user entity.User
	user, err = service.UserRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}
