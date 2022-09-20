package deco

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaak-ai/jaak-dabos/v1/transclude"
	"github.com/jaak-ai/jaak-japi/v1/armor"
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
