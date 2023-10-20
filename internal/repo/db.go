package repo

import (
	"gorm.io/gorm"
	"log"
	"users-services/internal/types"
)

type DbRepo struct {
	*gorm.DB
}

func New(db *gorm.DB) *DbRepo {
	return &DbRepo{db}
}

func (d DbRepo) MigrateModel() {
	err := d.AutoMigrate(&types.User{})
	if err != nil {
		log.Fatal("Error, migrate model: ", err)
	}
}
