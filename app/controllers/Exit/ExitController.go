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

// structを定義する
type ExitPUTReqBody struct {
	Id int `json:"id"`
	// デフォルトでもunknownが代入される
	Name string `json:"name"`
}

type ExitPUTResBody struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}
