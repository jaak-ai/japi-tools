package tracking

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${pid}] [${ip}] ${locals:requestid} ${status} - ${method} ${path}\n",
	})
}
