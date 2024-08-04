package request

type UserProfileUpdateRequest struct {
	UserId      int    `json:"user_id" validate:"required,numeric"`
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	BirthDate   string `json:"birth_date" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Address     string `json:"address" validate:"required"`
}
