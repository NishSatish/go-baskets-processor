package api

import (
	"github.com/gin-gonic/gin"
	v1 "go-basket-processor/pkg/service/v1"
	"net/http"
)

func getRouter(v1Service v1.V1ServiceInterface) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/baskets", func(c *gin.Context) {
		c.String(http.StatusOK, "donski")
	})

	v1Api := router.Group("/v1")
	{
		postGroup := v1Api.Group("/baskets")
		{
			postGroup.POST("/create", v1Service.CreateBasket)
		}
	}

	return router
}
