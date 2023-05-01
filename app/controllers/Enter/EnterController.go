package controllers

import "github.com/gin-gonic/gin"

// よりStrictに定義する
type EnterController struct{}

// httpのリクエストの一連の流れを、Contextが保持する
func (EnterController) EnterPut(c *gin.Context) {

	c.JSON(200, gin.H{
		"enter_time": "2020-01-01 00:00:00",
		"id":         12345,
		"name":       "DaikiMori",
	})
}
