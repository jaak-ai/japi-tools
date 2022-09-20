package processor

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

func New(options ...Options) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()
		if err != nil {
			return err
		}

		end := time.Now()
		duration := end.Sub(start)

		property := "duration"

		if len(options) > 0 && options[0].Property != "" {
			property = options[0].Property
		}

		bBody := ctx.Response().Body()
		mResponse := make(map[string]interface{})

		err = json.Unmarshal(bBody, &mResponse)
		if err != nil {
			return err
		}

		mResponse[property] = duration.String()

		bResponse, err := json.Marshal(mResponse)
		if err != nil {
			return err
		}

		ctx.Response().SetBody(bResponse)

		return nil
	}
}
