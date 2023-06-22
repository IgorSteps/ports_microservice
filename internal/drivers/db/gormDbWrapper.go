package db

import (
	"ports_microservice/internal/adapters/datastore"

	"gorm.io/gorm"
)

/*
	It usually considered bad practice to return an interface in Go, so why do I do it here?

	Using an interface like DBWrapper in this usecase can be a valid for returning an interface.
	It allows you to create a wrapper around the GORM database object and expose a more limited and
	controlled interface to the calling code.

	By returning DBWrapper from the methods in GormDBWrapper, I abstract away the specific implementation
	details of GORM and provide a consistent interface that can be used by other components in my application.
	This can be beneficial for decoupling code from GORM and making it easier to switch to a different database
	library or mock the database for testing purposes.

	Returning an interface also enables me to extend the functionality of the DBWrapper interface in the future, if needed,
	without modifying the calling code. This allows for better code maintainability and extensibility.

	In summary, using an interface like DBWrapper to wrap a database library like GORM is a common practice in Go.
	It provides a way to abstract the underlying implementation, promote code flexibility, and make it easier to test
	and switch database libraries if required.
*/

// GormDBWrapper implements datastore.DBWrapper interface.
type GormDBWrapper struct {
	db *gorm.DB
}

func NewGormDBWrapper(db *gorm.DB) *GormDBWrapper {
	return &GormDBWrapper{
		db: db,
	}
}

func (s *GormDBWrapper) Create(value interface{}) datastore.DBWrapper {
	return &GormDBWrapper{db: s.db.Create(value)}
}

func (s *GormDBWrapper) Error() error {
	return s.db.Error
}
