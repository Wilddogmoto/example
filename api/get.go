package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Number struct {
	Num   int    `json:"id"`
	Value string `json:"Value"`
}

var open = map[int]Number{
	1: {Num: 1, Value: "one"},
	2: {Num: 2, Value: "two"},
	3: {Num: 3, Value: "three"},
}

func getresponse() {
	router := gin.Default()

	router.GET("/data/:num", func(c *gin.Context) {

		var num Number
		if err := c.ShouldBindUri(&num); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	})

	fmt.Println()
	router.Run("localhost:8089")
}

func getreps() {

	router := gin.Default()

	router.GET("/data/:num", func(c *gin.Context) {

		num := c.Param("num")

		if num != "4" {
			c.JSON(http.StatusBadRequest, gin.H{"message is": "Error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message is": num})
	})

	router.Run("localhost:8089")
}
