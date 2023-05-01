package main

import (
	"github.com/DahuiSheng/gobackend/database"
	"github.com/DahuiSheng/gobackend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// host, portなど環境変数から読み込んだ方が良い
	client := database.GetClient()
	println(client)
	r := gin.Default()

	// gobackendにおけるrouterを参照
	// rは、router ginEngineに入れている
	router.Router(r)
	r.Run()
}
