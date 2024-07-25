package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"uaspw2/config"
	"uaspw2/helper"
	"uaspw2/models/web"
	"uaspw2/repositories"
)

type AuthService interface {
	Login(ctx context.Context, request web.LoginRequest) (string, error)
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

func (service *AuthServicesImpl) Login(ctx context.Context, request web.LoginRequest) (string, error) {
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

		return tokenString, nil
	} else {
		return "", errors.New("invalid username or password")
	}
}
