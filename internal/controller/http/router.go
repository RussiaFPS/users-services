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

	// Request Age API
	req, err := http.Get(types.AgeAPI + user.Name)
	if err != nil {
		log.Println("Error get age API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error read body age API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	userAgeApi := &types.UserAgeApi{}
	err = json.Unmarshal(body, &userAgeApi)
	if err != nil {
		log.Println("Error unmarshal body age API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	// Request Gender API
	req, err = http.Get(types.GenAPI + user.Name)
	if err != nil {
		log.Println("Error get gender API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	defer req.Body.Close()
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error read body gender API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	userGenderApi := &types.UserGenderApi{}
	err = json.Unmarshal(body, &userGenderApi)
	if err != nil {
		log.Println("Error unmarshal body gender API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	// Request Country API
	req, err = http.Get(types.CountryAPI + user.Name)
	if err != nil {
		log.Println("Error get country API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	defer req.Body.Close()
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error read body country API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	userCountryApi := &types.UserCountryApi{}
	err = json.Unmarshal(body, &userCountryApi)
	if err != nil {
		log.Println("Error unmarshal body country API: ", err)
		ctx.Status(http.StatusServiceUnavailable)
		return
	}

	if len(userCountryApi.Country) > 0 {
		u := types.User{Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: userAgeApi.Age,
			Gender: userGenderApi.Gender, CountryId: userCountryApi.Country[0].CountryId}
		ctx.JSON(201, u)
	} else {
		u := types.User{Name: user.Name, Surname: user.Surname, Patronymic: user.Patronymic, Age: userAgeApi.Age,
			Gender: userGenderApi.Gender, CountryId: ""}
		ctx.JSON(201, u)
	}
}
