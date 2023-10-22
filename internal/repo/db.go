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

func (d DbRepo) AddUser(user types.User) error {
	res := d.DB.Create(&user)
	if res.Error != nil {
		log.Println("Error, AddUser in repo: ", res.Error)
		return res.Error
	}

	return nil
}

func (d DbRepo) GetUser(id int) (types.User, error) {
	u := types.User{}
	res := d.DB.Where("id = ?", id).First(&u)

	if res.Error != nil {
		log.Println("Error, GetUser in repo: ", res.Error)
		return types.User{}, res.Error
	}

	return u, nil
}

func (d DbRepo) DelUser(id int) error {
	u := types.User{}
	res := d.DB.Where("id = ?", id).Delete(&u)

	if res.Error != nil {
		log.Println("Error, DelUser in repo: ", res.Error)
		return res.Error
	}

	return nil
}
