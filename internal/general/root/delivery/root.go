package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RootHandler struct {
	uc   usecase.Usecase
	conf *config.Config
	log  *logrus.Logger
}

func NewRootHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) RootHandler {
	return RootHandler{
		uc:   uc,
		conf: conf,
		log:  log,
	}
}

func (h RootHandler) GetRoot(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}