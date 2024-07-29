package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type ArticleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article
	Update(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article
	Delete(ctx context.Context, tx *sql.Tx, articleId int)
	FindByID(ctx context.Context, tx *sql.Tx, articleId int) (entity.Article, error)
	FindByUserID(ctx context.Context, tx *sql.Tx, userId int) []entity.Article
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Article
	FindAllByPublishStatus(ctx context.Context, tx *sql.Tx, publishStatus bool) []entity.Article
	FindByPublishStatusAndID(ctx context.Context, tx *sql.Tx, publishStatus bool, articleId int) (entity.Article, error)
	FindAllByPublishStatusAndUserID(ctx context.Context, tx *sql.Tx, publishStatus bool, userId int) []entity.Article
}

type ArticleRepositoryImpl struct {
}

func NewArticleRepository() ArticleRepository {
	return &ArticleRepositoryImpl{}
}

func (repository *ArticleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article {
	SQL := `INSERT INTO articles (user_id, title, description, is_published) VALUES (?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, article.UserId, article.Title, article.Description, article.IsPublished)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	article.Id = int(id)

	return article
}

func (repository *ArticleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article {
	SQL := `UPDATE articles SET title = ?, description = ?, is_published = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, article.Title, article.Description, article.IsPublished, article.Id)
	helper.PanicIfErr(err)

	return article
}

func (repository *ArticleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, articleId int) {
	SQL := `DELETE FROM articles WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, articleId)
	helper.PanicIfErr(err)
}

func (repository *ArticleRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, articleId int) (entity.Article, error) {
	SQL := `SELECT id, user_id, title, description, is_published, created_at, updated_at FROM articles WHERE id = ?`
	row, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfErr(err)
	defer row.Close()

	var article entity.Article

	if row.Next() {
		err := row.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt)
		helper.PanicIfErr(err)
		return article, nil
	} else {
		return article, errors.New("article not found")
	}
}

func (repository *ArticleRepositoryImpl) FindByUserID(ctx context.Context, tx *sql.Tx, userId int) []entity.Article {
	SQL := `SELECT id, user_id, title, description, is_published, created_at, updated_at FROM articles WHERE user_id = ?`
	row, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfErr(err)
	defer row.Close()

	var articles []entity.Article

	for row.Next() {
		var article entity.Article
		err := row.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt)
		articles = append(articles, article)
		helper.PanicIfErr(err)
	}

	return articles
}

func (repository *ArticleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Article {
	SQL := `SELECT id, user_id, title, description, is_published, created_at, updated_at FROM articles`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt)
		articles = append(articles, article)
		helper.PanicIfErr(err)
	}
	return articles
}

func (repository *ArticleRepositoryImpl) FindAllByPublishStatus(ctx context.Context, tx *sql.Tx, publishStatus bool) []entity.Article {
	SQL := `SELECT id, user_id, title, description, is_published, created_at, updated_at FROM articles WHERE is_published = ?`
	rows, err := tx.QueryContext(ctx, SQL, publishStatus)
	helper.PanicIfErr(err)
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt)
		articles = append(articles, article)
		helper.PanicIfErr(err)
	}
	return articles
}

func (repository *ArticleRepositoryImpl) FindByPublishStatusAndID(ctx context.Context, tx *sql.Tx, publishStatus bool, articleId int) (entity.Article, error) {
	SQL := `SELECT id, user_id, title, description, is_published, created_at, updated_at FROM articles WHERE is_published = ? AND id = ?`
	rows, err := tx.QueryContext(ctx, SQL, publishStatus, articleId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var article entity.Article

	if rows.Next() {
		err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt)
		helper.PanicIfErr(err)
		return article, nil
	}
	return article, errors.New("article not found")
}

func (repository *ArticleRepositoryImpl) FindAllByPublishStatusAndUserID(ctx context.Context, tx *sql.Tx, publishStatus bool, userId int) []entity.Article {
	SQL := `SELECT id, user_id, title, description, is_published, created_at, updated_at FROM articles WHERE is_published = ? AND user_id = ?`
	rows, err := tx.QueryContext(ctx, SQL, publishStatus, userId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt)
		articles = append(articles, article)
		helper.PanicIfErr(err)
	}
	return articles
}
