package request

type ArticleMediaCreateRequest struct {
	Type string `json:"type" validate:"required"`
	Path string `json:"path" validate:"required"`
}
