package request

type CommentRequest struct {
	UserId    int    `json:"user_id" validate:"required,numeric"`
	ArticleId int    `json:"article_id" validate:"required,numeric"`
	Comment   string `json:"comment" validate:"required"`
}
