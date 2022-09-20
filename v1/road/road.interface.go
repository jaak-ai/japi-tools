package road

import "github.com/gofiber/fiber/v2"

type RR interface {
	Add(method, path string, handlers ...fiber.Handler) fiber.Router
	Group(prefix string, handlers ...fiber.Handler) fiber.Router
}
