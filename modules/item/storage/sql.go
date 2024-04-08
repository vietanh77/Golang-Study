package storage

import "gorm.io/gorm"

type mySqlStore struct {
	db *gorm.DB
}

func NewMySqlStore(db *gorm.DB) *mySqlStore {
	return &mySqlStore{db: db}
}
