package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/models/entity"
	web2 "uaspw2/models/web"
	"uaspw2/repositories"
)

type UserService interface {
	Create(ctx context.Context, request web2.UserCreateRequest) web2.UserResponse
	Update(ctx context.Context, request web2.UserUpdateRequest) web2.UserResponse
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) web2.UserResponse
	FindAll(ctx context.Context) []web2.UserResponse
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

func (service *UserServiceImpl) Create(ctx context.Context, request web2.UserCreateRequest) web2.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	helper.PanicIfErr(err)

	user.Password = hashedPassword

	lastInsertID := service.UserRepository.Create(ctx, tx, user)

	user, _ = service.UserRepository.FindByID(ctx, tx, lastInsertID)
	return helper.ToUserResponse(user)

}

func (service *UserServiceImpl) Update(ctx context.Context, request web2.UserUpdateRequest) web2.UserResponse {
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

func (service *UserServiceImpl) FindByID(ctx context.Context, id int) web2.UserResponse {
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

func (service *UserServiceImpl) FindAll(ctx context.Context) []web2.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}
