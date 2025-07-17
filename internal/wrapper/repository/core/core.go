package core

import (
	"boilerplate/config"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
	user	"boilerplate/internal/core/user/repository"
)

type CoreRepository struct {
	User	user.Repository
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
		User:	user.NewUserRepo(dbList),
	}
}
