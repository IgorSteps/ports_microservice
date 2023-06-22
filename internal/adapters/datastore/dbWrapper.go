package datastore

// Implemented by db.GormDBWrapper.
type DBWrapper interface {
	Create(value interface{}) DBWrapper
	Error() error
}
