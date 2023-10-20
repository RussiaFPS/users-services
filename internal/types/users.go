package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	CountryId  string
}
