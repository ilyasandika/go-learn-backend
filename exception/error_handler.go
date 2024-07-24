package exception

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/web"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if notFoundError(c, err) {
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

func internalServerError(c *fiber.Ctx, err error) error {
	errorResponse := web.ErrorResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err.Error(),
	}
	return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
}
