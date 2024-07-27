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
	RegisterUser(ctx context.Context, tx *sql.Tx, request entity.User) entity.User
	CreateUserProfileOnRegisterUser(ctx context.Context, tx *sql.Tx, userId int)
}

type AuthRepositoryImpl struct {
}

func NewAuthenticationRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.User, error) {
	SQL := `SELECT id, username, password, role, created_at, updated_at FROM users WHERE username = ?`
	row, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfErr(err)
	defer row.Close()

	var user entity.User
	if row.Next() {
		err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfErr(err)
		return user, nil
	} else {
		return user, errors.New("user does not exist")
	}
}

func (repository *AuthRepositoryImpl) RegisterUser(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := `INSERT INTO users (username, password, role) VALUES (?, ?, ?)`

	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Role)
	helper.PanicIfErr(err)

	lastInsertId, err := result.LastInsertId()
	helper.PanicIfErr(err)

	user.Id = int(lastInsertId)

	return user
}

func (repository *AuthRepositoryImpl) CreateUserProfileOnRegisterUser(ctx context.Context, tx *sql.Tx, userId int) {
	SQL := `INSERT INTO user_profiles (user_id) VALUES (?)`
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfErr(err)
}
