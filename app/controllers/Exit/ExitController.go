package controllers

import "github.com/gin-gonic/gin"

// よりStrictに定義する
type ExitController struct{}

// httpのリクエストの一連の流れを、Contextが保持する
func (ExitController) ExitPost(c *gin.Context) {

	c.JSON(200, gin.H{
		"id":   "12345",
		"name": "DaikiMori",
	})
}
