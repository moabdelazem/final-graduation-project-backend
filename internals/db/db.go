package db

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(addr string) (*sql.DB, error) {
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db.DB()
}
