package deco

import (
	"ai.jaak.nfury/modules/armor"
	"ai.jaak.nfury/modules/transclude"
	"github.com/gofiber/fiber/v2"
)

func Popali[T any](factory func() T, ctx *fiber.Ctx) (T, error) {
	model := factory()

	err := transclude.PopulateFromRequest(model, ctx)
	if err != nil {
		return model, err
	}

	err = armor.ValidateStruct(model)
	if err != nil {
		return model, err
	}

	return model, nil
}
