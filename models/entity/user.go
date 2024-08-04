package entity

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserWithProfile struct {
	Id        int         `json:"id"`
	Username  string      `json:"username"`
	Profile   UserProfile `json:"profile"`
	Password  string      `json:"password"`
	Role      string      `json:"role"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}
