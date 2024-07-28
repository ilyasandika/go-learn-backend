package helper

import (
	"database/sql"
	"strconv"
	"uaspw2/exception"
	"uaspw2/models/entity"
	"uaspw2/models/web/response"
)

func ToUserResponse(user entity.User) response.UserResponse {
	return response.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []entity.User) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToUserProfileResponse(userProfile entity.UserProfile) response.UserProfileResponse {
	return response.UserProfileResponse{
		UserId:      userProfile.UserId,
		FullName:    userProfile.FullName,
		Gender:      userProfile.Gender,
		BirthDate:   userProfile.BirthDate,
		PhoneNumber: userProfile.PhoneNumber,
		Address:     userProfile.Address,
		CreatedAt:   userProfile.CreatedAt,
		UpdatedAt:   userProfile.UpdatedAt,
	}
}

func ToUserProfileResponses(userProfiles []entity.UserProfile) []response.UserProfileResponse {
	var userProfileResponses []response.UserProfileResponse
	for _, userProfile := range userProfiles {
		userProfileResponses = append(userProfileResponses, ToUserProfileResponse(userProfile))
	}
	return userProfileResponses
}

func ToUserPhotoProfileResponse(userPhotoProfile entity.UserProfilePhoto) response.UserProfilePhotoResponse {
	return response.UserProfilePhotoResponse{
		UserId:    userPhotoProfile.UserId,
		Path:      userPhotoProfile.Path,
		CreatedAt: userPhotoProfile.CreatedAt,
		UpdatedAt: userPhotoProfile.UpdatedAt,
	}
}

func ToIntFromParams(params string) int {
	id, err := strconv.Atoi(params)
	if err != nil {
		panic(exception.NewInvalidParameter("Invalid parameter. Excepted a number but received string"))
	}
	return id
}

func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
