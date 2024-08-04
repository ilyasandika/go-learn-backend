package response

import "uaspw2/models/entity"

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserWithProfileResponse struct {
	Id        int                `json:"id"`
	Username  string             `json:"username"`
	Profile   entity.UserProfile `json:"profile"`
	Role      string             `json:"role"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
}
