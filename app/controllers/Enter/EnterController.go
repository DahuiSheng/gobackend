package controllers

import "github.com/gin-gonic/gin"

// よりStrictに定義する
type EnterController struct{}

// httpのリクエストの一連の流れを、Contextが保持する
func (EnterController) EnterPut(c *gin.Context) {

	c.JSON(200, gin.H{
		"name": "DaikiMori",
	})
}

// structを定義する
type EnterPOSTReqBody struct {
	Name string
}
