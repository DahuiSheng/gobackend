package main

import (
	"context"
	"log"
	"net/http"

	"github.com/DahuiSheng/gobackend/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// host, portなど環境変数から読み込んだ方が良い
	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
