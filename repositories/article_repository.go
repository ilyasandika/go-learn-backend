package repositories

import (
	"context"
	"database/sql"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type ArticleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article
	CreateMedia(ctx context.Context, tx *sql.Tx, media entity.ArticleMedia) entity.ArticleMedia
	Update(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article
	Delete(ctx context.Context, tx *sql.Tx, articleId int)
	FindByID(ctx context.Context, tx *sql.Tx, articleId int) (entity.Article, error)
	FindAllByPublishStatus(ctx context.Context, tx *sql.Tx, publishStatus bool) []entity.Article
	FindAllByPublishStatusAndUserID(ctx context.Context, tx *sql.Tx, publishStatus bool, userId int) []entity.Article
	UpdatePublishStatus(ctx context.Context, tx *sql.Tx, articleId int, status bool)
}

type ArticleRepositoryImpl struct {
}

func NewArticleRepository() ArticleRepository {
	return &ArticleRepositoryImpl{}
}

func (repository *ArticleRepositoryImpl) UpdatePublishStatus(ctx context.Context, tx *sql.Tx, articleId int, status bool) {
	SQL := "UPDATE articles SET is_published = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, status, articleId)
	helper.PanicIfErr(err)
}

func (repository *ArticleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, article entity.Article) entity.Article {
	SQL := `INSERT INTO articles (user_id, title, description, content, is_published) VALUES (?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, article.UserId, article.Title, article.Description, article.Content, article.IsPublished)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	article.Id = int(id)

	return article
}

func (repository *ArticleRepositoryImpl) CreateMedia(ctx context.Context, tx *sql.Tx, media entity.ArticleMedia) entity.ArticleMedia {
	SQL := `INSERT INTO article_medias (article_id, type, path) VALUES (?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, media.ArticleId, media.Type, media.Path)
	helper.PanicIfErr(err)

	id, err := result.LastInsertId()
	media.Id = int(id)

	return media
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
	SQL := `SELECT 
				a.id,
				a.user_id,
				a.title,
				a.description,
				a.content,
				a.is_published,
				a.created_at,
				a.updated_at,
				up.full_name,
				am.id AS media_id,
				am.type AS media_type,
				am.path AS media_path
			FROM 
				articles a
			JOIN 
				user_profiles up ON a.user_id = up.user_id
			LEFT JOIN 
				article_medias am ON a.id = am.article_id
			WHERE 
				a.id = ?`
	rows, err := tx.QueryContext(ctx, SQL, articleId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var article entity.Article
	var media []entity.ArticleMedia

	for rows.Next() {
		var mediaID sql.NullInt64
		var mediaType sql.NullString
		var mediaPath sql.NullString

		err := rows.Scan(
			&article.Id,
			&article.UserId,
			&article.Title,
			&article.Description,
			&article.Content,
			&article.IsPublished,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.Author,
			&mediaID,
			&mediaType,
			&mediaPath,
		)
		if err != nil {
			return entity.Article{}, err
		}

		if mediaID.Valid {
			media = append(media, entity.ArticleMedia{
				Id:   int(mediaID.Int64),
				Type: mediaType.String,
				Path: mediaPath.String,
			})
		}
	}

	article.Media = media
	return article, nil
}

func (repository *ArticleRepositoryImpl) FindAllByPublishStatus(ctx context.Context, tx *sql.Tx, publishStatus bool) []entity.Article {
	SQL := `SELECT 
					a.id,
					a.user_id,
					a.title,
					a.description,
					a.content,
					a.is_published,
					a.created_at,
					a.updated_at,
					up.full_name
				FROM 
					articles a
				JOIN 
					user_profiles up ON a.user_id = up.user_id
				WHERE a.is_published = ?`
	rows, err := tx.QueryContext(ctx, SQL, publishStatus)
	helper.PanicIfErr(err)
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.Content, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt, &article.Author)
		articles = append(articles, article)
		helper.PanicIfErr(err)
	}
	return articles
}

func (repository *ArticleRepositoryImpl) FindAllByPublishStatusAndUserID(ctx context.Context, tx *sql.Tx, publishStatus bool, userId int) []entity.Article {
	SQL := `SELECT 
					a.id,
					a.user_id,
					a.title,
					a.description,
					a.content,
					a.is_published,
					a.created_at,
					a.updated_at,
					up.full_name
				FROM 
					articles a
				JOIN 
					user_profiles up ON a.user_id = up.user_id
				WHERE 
					a.is_published = ? 
					AND a.user_id = ?;
`
	rows, err := tx.QueryContext(ctx, SQL, publishStatus, userId)
	helper.PanicIfErr(err)
	defer rows.Close()

	var articles []entity.Article

	for rows.Next() {
		var article entity.Article
		err := rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Description, &article.Content, &article.IsPublished, &article.CreatedAt, &article.UpdatedAt, &article.Author)
		articles = append(articles, article)
		helper.PanicIfErr(err)
	}
	return articles
}
