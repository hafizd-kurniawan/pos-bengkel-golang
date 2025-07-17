package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NotFoundHandler struct {
	uc   usecase.Usecase
	conf *config.Config
	log  *logrus.Logger
}

func NewNotFoundHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) NotFoundHandler {
	return NotFoundHandler{
		uc:   uc,
		conf: conf,
		log:  log,
	}
}

func (h NotFoundHandler) GetNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Not Found",
	})
}