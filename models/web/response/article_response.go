package response

import "uaspw2/models/entity"

type ArticleResponse struct {
	Id          int                   `json:"id"`
	UserId      int                   `json:"user_id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Content     string                `json:"content"`
	Author      string                `json:"author"`
	IsPublished bool                  `json:"is_published"`
	Media       []entity.ArticleMedia `json:"media"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
}
