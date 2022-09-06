package catcher

import (
	"ai.jaak.nfury/usecase/tracer/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func New() fiber.ErrorHandler {
	return func(context *fiber.Ctx, err error) error {
		tracerI := context.Locals("tracer.enum")
		if tracerI != nil {
			tracer := tracerI.(repository.TracerRepository)
			_ = mgm.Coll(&tracer).Create(&tracer)
		}

		code := fiber.StatusInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return context.Status(code).JSON(HTTPError{
			Status:  code,
			Message: err.Error(),
		})
	}
}
