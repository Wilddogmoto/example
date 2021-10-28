package test

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pers struct {
	Value string `json:"value" binding:"required"`
}

func Postresponse() {

	router := gin.Default()

	router.POST("/:some", func(c *gin.Context) {

		var val pers

		if err := c.BindJSON(&val); err != nil {
			api.sendResponse(2, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": rev(val.Value)})
	})
	if err := router.Run("localhost:8089"); err != nil {
		fmt.Print("run error")
	}
}

func rev(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
