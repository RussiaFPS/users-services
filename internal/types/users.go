package types

import "gorm.io/gorm"

const AgeAPI = "https://api.agify.io/?name="
const GenAPI = "https://api.genderize.io/?name="
const CountryAPI = "https://api.nationalize.io/?name="

type User struct {
	gorm.Model
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	CountryId  string
}

type CountryInfo struct {
	CountryId   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type UsersReq struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UserAgeApi struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type UserGenderApi struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type UserCountryApi struct {
	Count   int           `json:"count"`
	Name    string        `json:"name"`
	Country []CountryInfo `json:"country"`
}

type UserAgeChan struct {
	Age *UserAgeApi
	Err error
}

type UserGenChan struct {
	Gen *UserGenderApi
	Err error
}

type UserCountryChan struct {
	Cnt *UserCountryApi
	Err error
}
