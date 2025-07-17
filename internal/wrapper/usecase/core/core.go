package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
	user	"boilerplate/internal/core/user/usecase"
)

type CoreUsecase struct {
	User	user.Usecase
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
		User:	user.NewUserUsecase(repo, conf, dbList, log),
	}
}
