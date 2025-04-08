package db

import (
	"todoProject/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *config.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Db{db}
}
