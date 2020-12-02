package dao

import "week02/internal/model"

func (d *Dao) GetUserList()([]*model.User,error) {
	user := model.User{}
	return user.List(d.engine)

}

