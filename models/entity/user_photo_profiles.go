package entity

type UserProfilePhoto struct {
	UserId    int    `json:"user_id"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
