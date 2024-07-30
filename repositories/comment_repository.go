package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type CommentRepository interface {
	Create(ctx context.Context, tx *sql.Tx, comment entity.Comment) entity.Comment
	FindByID(ctx context.Context, tx *sql.Tx, commentId int) (entity.Comment, error)
	FindByArticleID(ctx context.Context, tx *sql.Tx, articleId int) []entity.Comment
	Delete(ctx context.Context, tx *sql.Tx, commentId int, userId int)
}

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (c CommentRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, comment entity.Comment) entity.Comment {
	SQL := `INSERT INTO comments (user_id, article_id, comment) VALUES (?,?,?)`
	result, err := tx.ExecContext(ctx, SQL, comment.UserId, comment.ArticleId, comment.Comment)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	helper.PanicIfErr(err)

	comment.Id = int(id)

	return comment
}

func (c CommentRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, commentId int) (entity.Comment, error) {
	SQL := `SELECT id, user_id, article_id, comment, created_at, updated_at FROM comments WHERE id = ?`
	row, err := tx.QueryContext(ctx, SQL, commentId)
	helper.PanicIfErr(err)
	defer row.Close()

	var comment entity.Comment
	if row.Next() {
		err := row.Scan(&comment.Id, &comment.UserId, &comment.ArticleId, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt)
		helper.PanicIfErr(err)
		return comment, nil
	} else {
		return comment, errors.New("comment not found")
	}
}

func (c CommentRepositoryImpl) FindByArticleID(ctx context.Context, tx *sql.Tx, articleId int) []entity.Comment {
	SQL := `SELECT id, user_id, article_id, comment, created_at, updated_at FROM comments WHERE article_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.ArticleId, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt)
		helper.PanicIfErr(err)
		comments = append(comments, comment)
	}
	return comments
}

func (c CommentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, commentId int, userId int) {
	SQL := `DELETE FROM comments WHERE id = ? and user_id = ?`
	_, err := tx.ExecContext(ctx, SQL, commentId, userId)
	helper.PanicIfErr(err)
}
