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
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/repositories"
)

type AuthService interface {
	Login(ctx context.Context, request request.LoginRequest) string
	RegisterUser(ctx context.Context, request request.RegisterRequest) response.UserResponse
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

var TokenExpiresTime = time.Now().Add(time.Hour * 24)

func (service *AuthServicesImpl) Login(ctx context.Context, request request.LoginRequest) string {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.AuthRepository.GetUserByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(exception.NewInvalidCredentialsError("Invalid username or password"))
	}

	if helper.CheckPasswordHash(request.Password, user.Password) {
		//generate token
		claims := config.UserClaims{
			Id:       user.Id,
			Username: user.Username,
			Role:     user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(TokenExpiresTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(config.SecretKey)
		helper.PanicIfErr(err)

		return tokenString
	} else {
		panic(exception.NewInvalidCredentialsError("Invalid username or password"))
	}
}

func (service *AuthServicesImpl) RegisterUser(ctx context.Context, request request.RegisterRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	req := entity.User{
		Username: request.Username,
		Password: request.Password,
		Role:     "user",
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	helper.PanicIfErr(err)

	req.Password = hashedPassword

	userRegister := service.AuthRepository.RegisterUser(ctx, tx, req)
	req.Id = userRegister.Id

	userResponse, _ := service.AuthRepository.GetUserByUsername(ctx, tx, req.Username)

	service.AuthRepository.CreateUserProfileOnRegisterUser(ctx, tx, req.Id)

	defaultPhotoProfile := entity.UserProfilePhoto{
		UserId: req.Id,
		Path:   "default_profile_photo.svg",
	}
	service.AuthRepository.CreateUserPhotoProfileOnRegisterUser(ctx, tx, defaultPhotoProfile)

	return helper.ToUserResponse(userResponse)
}
