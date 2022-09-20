package bracer

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jaak-ai/jaak-japi/v1/tracer"
	"time"
)

func Middleware(optionsList ...OptionBracer) fiber.Handler {
	option := OptionBracer{
		IncludeRequest:  false,
		IncludeResponse: false,
	}

	if len(optionsList) > 0 {
		option = optionsList[0]
	}

	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		bcr := New()
		bcr.SetEventId(uuid.New().String())
		bcr.SetRequest(&tracer.RequestTracer{
			Ip:     ctx.IP(),
			Method: tracer.MapperMethodRequest(ctx.Method()),
			Path:   ctx.Path(),
			Meta:   &tracer.MetaRequestTracer{},
		})

		requestIdHeader := ctx.Get("Request-Id")
		if requestIdHeader != "" {
			bcr.Request.SetId(requestIdHeader)
		}

		if option.IncludeRequest {
			bcr.Request.Meta.Request, _ = mapperBytesByJson(ctx.Request().Body())
		}

		ctx.Locals("bracer", bcr)

		err := ctx.Next()

		end := time.Now()
		duration := end.Sub(start)
		bcr.ProcessTime = duration.Seconds()

		if option.IncludeResponse {
			bcr.Request.Meta.Response, _ = mapperBytesByJson(ctx.Response().Body())
		}

		if err != nil {
			bcr.SetMessage(err.Error())

			if e, ok := err.(*fiber.Error); ok {
				bcr.Request.SetStatusCode(e.Code)
			} else {
				bcr.Request.SetStatusCode(fiber.StatusInternalServerError)
			}
		} else {
			bcr.Request.SetStatusCode(ctx.Response().StatusCode())
		}

		_ = bcr.Save()

		return err
	}
}

func mapperBytesByJson(b []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
