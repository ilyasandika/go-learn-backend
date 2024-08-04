package entity

type ArticleMedia struct {
	Id        int    `json:"id"`
	ArticleId int    `json:"article_id"`
	Type      string `json:"type"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
