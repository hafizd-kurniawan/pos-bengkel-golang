package repository

import (
	"boilerplate/pkg/infra/db"
)

type Repository interface {
}

type UserRepo struct {
	DBList *db.DatabaseList
}

func NewUserRepo(dbList *db.DatabaseList) UserRepo {
	return UserRepo{
		DBList: dbList,
	}
}