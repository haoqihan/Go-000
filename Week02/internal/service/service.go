package service

import (
	"week02/internal/dao"
	"week02/internal/db"
)

type Service struct {
	dao *dao.Dao
}

func New() Service {
	return Service{dao: dao.New(db.New())}
}