package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// よりStrictに定義する
type EnterController struct{}

// httpのリクエストの一連の流れを、Contextが保持する
func (EnterController) EnterPut(c *gin.Context) {

	// :=は、代入と型定義を一緒に行なっている
	json := EnterPUTReqBody{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resJson := EnterPUTResBody{
		Id:   12345,
		Name: *json.Name,
	}

	// JSONを返す、200はステータスコード
	c.JSON(200, resJson)
}

// structを定義する
type EnterPUTReqBody struct {
	// 必須でないものは、ポインタとして扱う：何も入ってない時は何も入らない
	Name *string `json:"name"`
}

type EnterPUTResBody struct {
	// nameをunknownでも返す
	// *stringとstringは、型が違うのでだめ。
	Id   int    `json:"id"`
	Name string `json:"name"`
}
