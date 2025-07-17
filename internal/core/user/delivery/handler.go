package delivery

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewUserHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) UserHandler {
	return UserHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}