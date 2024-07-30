package request

type LikeRequest struct {
	UserId    int `json:"user_id" validate:"required,numeric"`
	ArticleId int `json:"article_id" validate:"required,numeric"`
}
