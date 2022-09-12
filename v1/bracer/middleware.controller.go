package bracer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jaak-ai/jaak-japi/v1/tracer"
	"time"
)

func Middleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		bcr := New()
		bcr.SetRequest(&tracer.RequestTracer{
			Ip:     ctx.IP(),
			Method: tracer.MapperMethodRequest(ctx.Method()),
			Path:   ctx.Path(),
		})
		ctx.Locals("bracer", bcr)

		err := ctx.Next()

		end := time.Now()
		duration := end.Sub(start)
		bcr.ProcessTime = duration.Seconds()

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
