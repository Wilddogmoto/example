package api

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func queryUser(c *gin.Context) {

	var val *data.Users

	if err := c.BindJSON(&val); err != nil {
		sendResponse(2, c)
		return
	}

	queryId(val, data.DataBase)

	sendResponse(1, c)

}

func queryId(j *data.Users, db *gorm.DB) {

	if err := db.Model(&data.Users{}).Where(j, true).Update("name", "Tom").Error; err != nil {
		fmt.Println("query error", err)
		return
	}
}
