package main

import (
	"net/http"

	"github.com/DahuiSheng/gobackend/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// host, portなど環境変数から読み込んだ方が良い
	client := database.GetClient()
	println(client)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
