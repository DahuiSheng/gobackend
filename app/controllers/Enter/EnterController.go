package controllers

import "github.com/gin-gonic/gin"

// よりStrictに定義する
type EnterController struct{}

// httpのリクエストの一連の流れを、Contextが保持する
func (EnterController) EnterPut(c *gin.Context) {

	// :=は、代入と型定義を一緒に行なっている
	json := EnterPOSTReqBody{}

	c.JSON(200, gin.H{
		"name": "DaikiMori",
	})
}

// structを定義する
type EnterPOSTReqBody struct {
	// 必須でないものは、ポインタとして扱う
	Name *string `json:"name"`
}

type EnterPOSTResBody struct {
	Id int `json:"id"`
}
