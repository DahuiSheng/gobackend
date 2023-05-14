package router

import (
	EnterController "github.com/DahuiSheng/gobackend/app/controllers/Enter"
	ExitController "github.com/DahuiSheng/gobackend/app/controllers/Exit"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	EnterController := EnterController.EnterController{}
	router.PUT("/enter", EnterController.EnterPut)

	ExitController := ExitController.ExitController{}
	router.POST("/exit", ExitController.ExitPost)
}

type ExitPUTResBody struct {
	Name
}
