package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
		api.GET("/user", c.get)
		api.GET("/user/:id", c.getById)
		api.PUT("/update/:id", c.update)
		api.DELETE("/user/:id", c.delById)
	}
}

func (c Controller) update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		log.Println("Error, update conv to int: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Update read request body err: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	user := types.UsersReq{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		log.Println("Update unmarshal request body err: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	if user.Name == "" || user.Surname == "" {
		log.Println("Update empty fields: name or surname")
		ctx.Status(http.StatusBadRequest)
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
		u := types.User{Model: gorm.Model{ID: uint(id)}, Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: chAgeStr.Age.Age,
			Gender: chGenStr.Gen.Gender, CountryId: chCntStr.Cnt.Country[0].CountryId}

		if c.Db.UpdateUser(u) != nil {
			ctx.Status(http.StatusBadGateway)
			return
		}

		ctx.JSON(200, u)
	} else {
		u := types.User{Model: gorm.Model{ID: uint(id)}, Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: chAgeStr.Age.Age,
			Gender: chGenStr.Gen.Gender, CountryId: ""}

		if c.Db.UpdateUser(u) != nil {
			ctx.Status(http.StatusBadGateway)
			return
		}

		ctx.JSON(200, u)
	}
}

func (c Controller) get(ctx *gin.Context) {
	res, err := c.Db.GetUser()
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(200, res)
}

func (c Controller) delById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		log.Println("Error, delById conv to int: ", err)
		ctx.Status(http.StatusNotFound)
		return
	}

	err = c.Db.DelUser(id)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(200)
}

func (c Controller) getById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 0 {
		log.Println("Error, getById conv to int: ", err)
		ctx.Status(http.StatusNotFound)
		return
	}

	res, err := c.Db.GetUserById(id)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(200, res)
}

func (c Controller) add(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Add read request body err: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	user := types.UsersReq{}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		log.Println("Add unmarshal request body err: ", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	if user.Name == "" || user.Surname == "" {
		log.Println("Add empty fields: name or surname")
		ctx.Status(http.StatusBadRequest)
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

		if c.Db.AddUser(u) != nil {
			ctx.Status(http.StatusBadGateway)
			return
		}

		ctx.JSON(201, u)
	} else {
		u := types.User{Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: chAgeStr.Age.Age,
			Gender: chGenStr.Gen.Gender, CountryId: ""}

		if c.Db.AddUser(u) != nil {
			ctx.Status(http.StatusBadGateway)
			return
		}

		ctx.JSON(201, u)
	}
}
