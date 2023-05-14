package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// よりStrictに定義する
type ExitController struct{}

// httpのリクエストの一連の流れを、Contextが保持する
func (ExitController) ExitPost(c *gin.Context) {

	json := ExitPUTReqBody{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resJson := ExitPUTResBody{
		Id:       json.Id,
		Name:     json.Name,
		CheckIn:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
		CheckOut: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	}

	c.JSON(200, resJson)
}

// structを定義する
type ExitPUTReqBody struct {
	Id int `json:"id" binding:"required"`
	// デフォルトでもunknownが代入される
	Name string `json:"name"`
}

type ExitPUTResBody struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
}
