package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"uaspw2/models/web"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if notFoundError(c, err) {
		return nil
	}

	if validatorError(c, err) {
		return nil
	}

	return internalServerError(c, err)
}

func notFoundError(c *fiber.Ctx, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		errorResponse := web.ErrorResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Error:  exception.Error(),
		}
		return c.Status(fiber.StatusNotFound).JSON(errorResponse) == nil // c.Status... itu return nya error
	} else {
		return false
	}
}

func invalidCredentialError(c *fiber.Ctx, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		errorResponse := web.ErrorResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Error:  exception.Error(),
		}
		return c.Status(fiber.StatusNotFound).JSON(errorResponse) == nil // c.Status... itu return nya error
	} else {
		return false
	}
}

func validatorError(c *fiber.Ctx, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	var errorMessages []string
	for _, err := range exception {
		errorMessages = append(errorMessages, err.Translate(nil))
	}

	if ok {
		errorResponse := web.ErrorResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Error:  errorMessages,
		}
		return c.Status(fiber.StatusNotFound).JSON(errorResponse) == nil // c.Status... itu return nya error
	} else {
		return false
	}
}

func internalServerError(c *fiber.Ctx, err error) error {
	errorResponse := web.ErrorResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err.Error(),
	}
	return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
}
