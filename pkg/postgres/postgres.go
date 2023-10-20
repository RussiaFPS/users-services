package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getConfigurationPostgres()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error, connect to postgres: ", err)
	}

	return db
}
