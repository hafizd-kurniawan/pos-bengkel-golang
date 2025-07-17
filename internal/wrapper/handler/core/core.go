package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
	user	"boilerplate/internal/core/user/delivery"
)

type CoreHandler struct {
	User	user.UserHandler
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
		User:	user.NewUserHandler(uc, conf, log),
	}
}
