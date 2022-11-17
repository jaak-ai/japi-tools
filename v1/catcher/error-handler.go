package catcher

import (
	"github.com/gofiber/fiber/v2"
)

func New() fiber.ErrorHandler {
	return func(context *fiber.Ctx, err error) error {

		errorCatcher, ok := err.(*Error)
		if ok {
			return context.Status(errorCatcher.Status).JSON(err)
		}

		errorFiber, ok := err.(*fiber.Error)
		if ok {
			return context.Status(errorFiber.Code).JSON(Error{
				Status:  errorFiber.Code,
				Message: err.Error(),
			})
		}

		return context.
			Status(fiber.StatusInternalServerError).
			JSON(Error{
				Status:  fiber.StatusInternalServerError,
				Message: err.Error(),
			})
	}
}
