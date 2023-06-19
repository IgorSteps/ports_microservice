package datastore

import "database/sql"

type GormDBWrapper interface {
	Close() error
	DB() *sql.DB
	Create(value interface{}) error
}
