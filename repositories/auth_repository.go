package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
	"uaspw2/models/web"
)

type AuthRepository interface {
	GetPasswordByUsername(ctx context.Context, tx *sql.Tx, username string) (string, error)
	Login(ctx context.Context, tx *sql.Tx, request web.LoginRequest) (entity.User, error)
}

type AuthRepositoryImpl struct {
}

func NewAuthenticationRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetPasswordByUsername(ctx context.Context, tx *sql.Tx, username string) (string, error) {
	SQL := `SELECT password FROM users WHERE username = ?`
	row, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfErr(err)
	defer row.Close()

	var password string
	if row.Next() {
		err = row.Scan(&password)
		helper.PanicIfErr(err)
		return password, nil
	} else {
		return "", errors.New("user not found")
	}

}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, request web.LoginRequest) (entity.User, error) {
	SQL := `SELECT id, username, role FROM users WHERE username = ? LIMIT 1`
	row, err := tx.QueryContext(ctx, SQL, request.Username)
	helper.PanicIfErr(err)
	defer row.Close()

	var user entity.User
	if row.Next() {
		err := row.Scan(&user.Id, &user.Username, &user.Role)
		helper.PanicIfErr(err)
		return user, nil
	} else {
		return user, errors.New("unauthorized")
	}
}
