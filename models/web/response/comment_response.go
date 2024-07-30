package response

type CommentResponse struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ArticleId int    `json:"article_id"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
