package model

import (
	"github.com/pkg/errors"
	"week02/internal/db"
)

type User struct {
	Name string
	Age  string
}

func (u User) List(db *db.DB) ([]*User, error) {
	var list []*User
	err := db.Select("select * from user")
	if err != nil {
		return nil, errors.Wrap(err,"user.List error")
	}
	return list, nil
}
