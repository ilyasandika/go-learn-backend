package request

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=16"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=6,max=16"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Role     string `json:"role"`
}
