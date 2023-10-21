package http

import (
	"github.com/gin-gonic/gin"
	"users-services/internal/repo"
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
	//TODO add
}
