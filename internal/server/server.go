package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igromancer/tiny-ledger/internal/config"
)

var conf *config.Config

func init() {
	conf = config.Get()
	fmt.Println(conf)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Create() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", pingHandler)

	return router
}
