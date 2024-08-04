package request

type ArticleCreateRequest struct {
	UserId      int                         `json:"user_id" validate:"required,numeric"`
	Title       string                      `json:"title" validate:"required,max=255"`
	Description string                      `json:"description"`
	Content     string                      `json:"content"`
	IsPublished bool                        `json:"is_published" validate:"boolean"`
	Media       []ArticleMediaCreateRequest `json:"media"`
}

type ArticleUpdateRequest struct {
	Id          int    `json:"id" validate:"required,numeric"`
	UserId      int    `json:"user_id" validate:"required,numeric"`
	Title       string `json:"title" validate:"required,max=255"`
	Content     string `json:"content" validate:"required"`
	Description string `json:"description"`
	IsPublished bool   `json:"is_published" validate:"boolean"`
}
