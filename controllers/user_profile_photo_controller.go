package controllers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"uaspw2/helper"
	"uaspw2/models/web/request"
	"uaspw2/models/web/response"
	"uaspw2/services"
)

type UserProfilePhotoController interface {
	UpdateByToken(c *fiber.Ctx) error
	FindByToken(c *fiber.Ctx) error
}

type UserProfilePhotoControllerImpl struct {
	UserProfilePhotoService services.UserProfilePhotoService
}

func NewUserProfilePhotoController(userProfilePhotoService services.UserProfilePhotoService) UserProfilePhotoController {
	return &UserProfilePhotoControllerImpl{
		UserProfilePhotoService: userProfilePhotoService,
	}
}

func isValidProfilePhoto(file *multipart.FileHeader) (bool, error) {
	allowedTypes := map[string]bool{
		"image/jpeg":    true,
		"image/png":     true,
		"image/gif":     true,
		"image/svg+xml": true,
	}

	// Check file type
	fileType := file.Header.Get("Content-Type")
	if !allowedTypes[fileType] {
		return false, errors.New("invalid file type (JPEG, PNG, GIF or SVG only)")
	}

	// Check file size (example: limit to 2MB)
	const maxFileSize = 2 * 1024 * 1024 // 2MB
	if file.Size > maxFileSize {
		return false, errors.New("invalid file type (max 2 MB)")
	}

	return true, nil
}

func (controller *UserProfilePhotoControllerImpl) UpdateByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	file, err := c.FormFile("profilePhoto")
	helper.PanicIfErr(err)

	_, err = isValidProfilePhoto(file)
	helper.PanicIfErr(err)

	profile := controller.UserProfilePhotoService.FindByUserId(c.Context(), user.Id)

	oldFilePath := filepath.Join("./public/profile_photos", profile.Path)

	fileName := fmt.Sprintf("%d_%d%s", user.Id, time.Now().Unix(), filepath.Ext(file.Filename))
	filePath := fmt.Sprintf("./public/profile_photos/%s", fileName)

	err = c.SaveFile(file, filePath)
	helper.PanicIfErr(err)

	userProfilePhoto := request.UserProfilePhotoRequest{
		UserId: user.Id,
		Path:   fileName,
	}
	result := controller.UserProfilePhotoService.UpdateByUserId(c.Context(), userProfilePhoto)

	if profile.Path != "default_profile_photo.svg" {
		err = os.Remove(oldFilePath)
		helper.PanicIfErr(err)
	}

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "user profile photo updated successfully",
		Data:    result,
	}
	return c.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *UserProfilePhotoControllerImpl) FindByToken(c *fiber.Ctx) error {
	user, err := helper.GetUserByToken(c)
	helper.PanicIfErr(err)

	profile := controller.UserProfilePhotoService.FindByUserId(c.Context(), user.Id)

	webResponse := response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "user profile photo found",
		Data:    profile,
	}

	return c.Status(fiber.StatusOK).JSON(webResponse)
}
