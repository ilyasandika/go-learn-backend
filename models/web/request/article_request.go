package request

type ArticleCreateRequest struct {
	UserId      int    `json:"user_id" validate:"required,numeric"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
	IsPublished bool   `json:"is_published" validate:"boolean"`
}

type ArticleUpdateRequest struct {
	Id          int    `json:"id" validate:"required,numeric"`
	UserId      int    `json:"user_id" validate:"required,numeric"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
	IsPublished bool   `json:"is_published" validate:"boolean"`
}
