package web

type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=6,max=16"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role"`
}

type UserUpdateRequest struct {
	Id       int    `json:"id" validate:"required,numeric,gte=1"`
	Username string `json:"username" validate:"required,max=16,min=6"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=16"`
	Password string `json:"password" validate:"omitempty,min=6"`
}
