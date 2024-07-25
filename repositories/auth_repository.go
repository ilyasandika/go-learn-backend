package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error)
}

type AuthRepositoryImpl struct {
}

func NewAuthenticationRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error) {
	SQL := `SELECT id, username, password, role FROM users WHERE username = ?`
	row, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfErr(err)
	defer row.Close()

	var user entity.User
	if row.Next() {
		err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
		helper.PanicIfErr(err)
		return user, nil
	} else {
		return user, errors.New("user does not exist")
	}
}
