package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"users-services/internal/repo"
	"users-services/internal/types"
)

type Controller struct {
	Route *gin.Engine
	Db    *repo.DbRepo
}

func New(repo *repo.DbRepo) *Controller {
	return &Controller{gin.Default(), repo}
}

func (c Controller) NewRoute() {
	api := c.Route.Group("/api")
	{
		api.POST("/add", c.add)
	}
}

func (c Controller) add(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Add read request body err: ", err)
		return
	}
	user := types.UsersReq{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		log.Println("Add unmarshal request body err: ", err)
		return
	}
	if user.Name == "" || user.Surname == "" {
		log.Println("Add empty fields: name or surname")
		ctx.Status(http.StatusNotFound)
		return
	}

	// Create chan from requests
	chAge := make(chan types.UserAgeChan)
	chAgeStr := types.UserAgeChan{}
	chGen := make(chan types.UserGenChan)
	chGenStr := types.UserGenChan{}
	chCnt := make(chan types.UserCountryChan)
	chCntStr := types.UserCountryChan{}

	// Request Age API
	go RequestAge(user.Name, chAge)
	// Request Gender API
	go RequestGender(user.Name, chGen)
	// Request Country API
	go RequestCountry(user.Name, chCnt)

	// Get data requests
	chAgeStr = <-chAge
	chGenStr = <-chGen
	chCntStr = <-chCnt

	if chAgeStr.Err != nil || chGenStr.Err != nil || chCntStr.Err != nil {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	if len(chCntStr.Cnt.Country) > 0 {
		u := types.User{Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: chAgeStr.Age.Age,
			Gender: chGenStr.Gen.Gender, CountryId: chCntStr.Cnt.Country[0].CountryId}
		ctx.JSON(201, u)
	} else {
		u := types.User{Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: chAgeStr.Age.Age,
			Gender: chGenStr.Gen.Gender, CountryId: ""}
		ctx.JSON(201, u)
	}
}
