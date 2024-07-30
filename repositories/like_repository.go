package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type LikeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, like entity.Like) entity.Like
	FindByArticleAndUser(ctx context.Context, tx *sql.Tx, articleId int, userId int) (entity.Like, error)
	FindByArticleId(ctx context.Context, tx *sql.Tx, articleId int) []entity.Like
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []entity.Like
	Delete(ctx context.Context, tx *sql.Tx, articleId int, userId int)
}

type LikeRepositoryImpl struct {
}

func NewLikeRepository() LikeRepository {
	return &LikeRepositoryImpl{}
}

func (repository *LikeRepositoryImpl) FindByArticleAndUser(ctx context.Context, tx *sql.Tx, articleId int, userId int) (entity.Like, error) {
	SQL := `SELECT id, user_id, article_id, created_at, updated_at FROM likes WHERE article_id = ? AND user_id = ?`
	row, err := tx.QueryContext(ctx, SQL, articleId, userId)
	helper.PanicIfErr(err)
	defer row.Close()

	var like entity.Like
	if row.Next() {
		err := row.Scan(&like.Id, &like.UserId, &like.ArticleId, &like.CreatedAt, &like.UpdatedAt)
		helper.PanicIfErr(err)
		return like, nil
	} else {
		return like, errors.New("like not found")
	}

}

func (repository *LikeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, like entity.Like) entity.Like {
	SQL := "INSERT INTO likes (user_id, article_id) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, like.UserId, like.ArticleId)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	like.Id = int(id)

	return like
}

func (repository *LikeRepositoryImpl) FindByArticleId(ctx context.Context, tx *sql.Tx, articleId int) []entity.Like {
	SQL := "SELECT id, article_id, user_id, created_at, updated_at FROM likes WHERE article_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var likes []entity.Like

	for rows.Next() {
		var like entity.Like
		err := rows.Scan(&like.Id, &like.ArticleId, &like.UserId, &like.CreatedAt, &like.UpdatedAt)
		helper.PanicIfErr(err)
		likes = append(likes, like)
	}

	return likes
}

func (repository *LikeRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []entity.Like {
	SQL := "SELECT id, article_id, user_id, created_at, updated_at FROM likes WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var likes []entity.Like

	for rows.Next() {
		var like entity.Like
		err := rows.Scan(&like.Id, &like.ArticleId, &like.UserId, &like.CreatedAt, &like.UpdatedAt)
		helper.PanicIfErr(err)
		likes = append(likes, like)
	}

	return likes
}

func (repository *LikeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, articleId int, userId int) {
	SQL := `DELETE FROM likes WHERE article_id = ? AND user_id = ?`
	_, err := tx.ExecContext(ctx, SQL, articleId, userId)
	helper.PanicIfErr(err)
}
