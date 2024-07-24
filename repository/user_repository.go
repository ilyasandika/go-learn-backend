package repository

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/model/entity"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.User) int
	Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entity.User) int {
	SQL := `INSERT INTO users (username, password, role) VALUES (?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, &user.Username, &user.Password, &user.Role)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	return int(id)
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := `UPDATE users set username = ?, password = ?, role = ? where id = ?`
	_, err := tx.ExecContext(ctx, SQL, &user.Username, &user.Password, &user.Role, &user.Id)
	helper.PanicIfErr(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `DELETE FROM users WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (entity.User, error) {
	SQL := `SELECT id, username, password, role, created_at, updated_at FROM users WHERE id = ?`
	row, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer row.Close()

	user := entity.User{}
	if row.Next() {
		err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfErr(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.User {
	SQL := `SELECT id, username, role, created_at, updated_at FROM users`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		user := entity.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfErr(err)
		users = append(users, user)
	}
	return users
}
