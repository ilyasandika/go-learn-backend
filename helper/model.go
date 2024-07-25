package helper

import (
	"strconv"
	"uaspw2/models/entity"
	"uaspw2/models/web"
)

func ToUserResponse(user entity.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []entity.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToIntFromParams(params string) int {
	id, err := strconv.Atoi(params)
	PanicIfErr(err)
	return id
}
