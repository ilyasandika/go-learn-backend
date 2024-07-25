package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"uaspw2/config"
	"uaspw2/exception"
	"uaspw2/helper"
	"uaspw2/models/entity"
	"uaspw2/models/web"
	"uaspw2/repositories"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) string
	RegisterUser(ctx context.Context, request web.RegisterRequest) web.UserResponse
}

type AuthServicesImpl struct {
	AuthRepository repositories.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthenticationServices(authRepository repositories.AuthRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServicesImpl{
		AuthRepository: authRepository,
		DB:             db,
		Validate:       validate,
	}
}

var ExpiresTime = time.Now().Add(time.Hour * 24)

func (service *AuthServicesImpl) Login(ctx context.Context, request web.LoginRequest) string {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.AuthRepository.GetUserByUsername(ctx, tx, request.Username)
	helper.PanicIfErr(err)

	if helper.CheckPasswordHash(request.Password, user.Password) {
		//generate token
		claims := config.UserClaims{
			Id:       user.Id,
			Username: user.Username,
			Role:     user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(ExpiresTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(config.SecretKey)
		helper.PanicIfErr(err)

		return tokenString
	} else {
		panic(exception.NewInvalidCredentialsError("invalid username or password"))
	}
}

func (service *AuthServicesImpl) RegisterUser(ctx context.Context, request web.RegisterRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user := entity.User{
		Username: request.Username,
		Password: request.Password,
		Role:     "user",
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	helper.PanicIfErr(err)

	user.Password = hashedPassword

	user = service.AuthRepository.RegisterUser(ctx, tx, user)

	user, _ = service.AuthRepository.GetUserByUsername(ctx, tx, user.Username)
	return helper.ToUserResponse(user)
}
