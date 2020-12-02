package db

import "database/sql"

type DB struct {
}

func (db *DB) Select(i interface{}) error {
	return sql.ErrNoRows
}

func New() *DB {
	return &DB{}
}
