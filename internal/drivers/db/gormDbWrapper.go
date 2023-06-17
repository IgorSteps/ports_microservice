package db

import (
	"gorm.io/gorm"
)

type GormDBWrapper struct {
	w *gorm.DB
}

func NewGormDBWrapper(w *gorm.DB) *GormDBWrapper {
	return &GormDBWrapper{w}
}

func (s *GormDBWrapper) Create(value interface{}) error {
	return s.w.Create(value).Error
}
