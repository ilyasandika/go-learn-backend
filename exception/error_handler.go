package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"uaspw2/models/web/response"
)

func ErrorHandler(c *fiber.Ctx, err error) error {

	if notFoundError(c, err) {
		return nil
	}

	if validatorError(c, err) {
		return nil
	}

	if invalidCredentialError(c, err) {
		return nil
	}

	if invalidParameterError(c, err) {
		return nil
	}

	return internalServerError(c, err)
}

func notFoundError(c *fiber.Ctx, err error) bool {
	var exception *NotFoundError
	if errors.As(err, &exception) {
		errorResponse := response.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "NOT FOUND",
			Error:   exception.Error(),
		}
		return c.Status(fiber.StatusNotFound).JSON(errorResponse) == nil // c.Message... itu return nya error
	} else {
		return false
	}
}

func invalidCredentialError(c *fiber.Ctx, err error) bool {
	var exception *InvalidCredentialsError
	if errors.As(err, &exception) {
		errorResponse := response.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "UNAUTHORIZED",
			Error:   exception.Error(),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(errorResponse) == nil // c.Message... itu return nya error
	} else {
		return false
	}
}

func invalidParameterError(c *fiber.Ctx, err error) bool {
	var exception *InvalidParameterError
	if errors.As(err, &exception) {
		errorResponse := response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "INVALID PARAMETER",
			Error:   exception.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse) == nil // c.Message... itu return nya error
	} else {
		return false
	}
}

func validatorError(c *fiber.Ctx, err error) bool {
	var exception validator.ValidationErrors

	var errorMessages []string
	for _, err := range exception {
		errorMessages = append(errorMessages, err.Translate(nil))
	}

	if errors.As(err, &exception) {
		errorResponse := response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   errorMessages,
		}
		return c.Status(fiber.StatusNotFound).JSON(errorResponse) == nil // c.Message... itu return nya error
	} else {
		return false
	}
}

func internalServerError(c *fiber.Ctx, err error) error {
	errorResponse := response.ErrorResponse{
		Code:    fiber.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Error:   err.Error(),
	}
	return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
}
