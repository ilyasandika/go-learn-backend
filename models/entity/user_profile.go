package entity

type UserProfile struct {
	UserId      int    `json:"user_id"`
	FullName    string `json:"full_name"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
