package http

import (
	"github.com/gin-gonic/gin"
	"users-services/internal/repo"
)

type Controller struct {
	route *gin.Engine
	db    *repo.DbRepo
}

func New(repo *repo.DbRepo) *gin.Engine {
	r := gin.Default()
	c := &Controller{r, repo}

	api := c.route.Group("/api")
	{
		api.POST("/add", c.add)
	}

	return r
}

func (c Controller) add(ctx *gin.Context) {
	//TODO add
}
