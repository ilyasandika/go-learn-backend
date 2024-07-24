package services

import (
	"context"
	"database/sql"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/model/entity"
	"uaspw2/repository"
	"uaspw2/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, id int)
	FindByID(ctx context.Context, id int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}
	lastInsertID := service.UserRepository.Create(ctx, tx, user)

	user, _ = service.UserRepository.FindByID(ctx, tx, lastInsertID)
	return helper.ToUserResponse(user)

}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	var user entity.User
	user, err = service.UserRepository.FindByID(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Username = request.Username
	user.Password = request.Password
	user.Role = request.Role

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

func (service *UserServiceImpl) FindByID(ctx context.Context, id int) web.UserResponse {
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

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}
