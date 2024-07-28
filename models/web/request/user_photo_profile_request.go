package request

type UserProfilePhotoRequest struct {
	UserId int    `json:"user_id" validate:"required,numeric"`
	Path   string `json:"path" validate:"required"`
}
