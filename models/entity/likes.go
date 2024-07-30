package entity

type Like struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ArticleId int    `json:"article_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
