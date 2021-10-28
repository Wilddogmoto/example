package test

import (
	"github.com/Wilddogmoto/example_project/api"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
)

func comparePassword(c *gin.Context) {

	var val *data.Users

	if err := c.BindJSON(&val); err != nil {
		api.sendResponse(2, c)
		return
	}

	api.sendResponse(1, c)
}

func compareId() {

}
