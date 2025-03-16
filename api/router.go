package api

import (
	"github.com/gin-gonic/gin"
	"go-basket-processor/pkg/db/models/Basket"
	"net/http"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/baskets", func(c *gin.Context) {
		Basket.CreateTable()
		c.String(http.StatusOK, "donski")
	})

	return router
}
