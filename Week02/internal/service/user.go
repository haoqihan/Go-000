package service

import (
	"week02/internal/model"
)

func (svc *Service) GetUserList() ([]*model.User, error) {
	return svc.dao.GetUserList()
}
