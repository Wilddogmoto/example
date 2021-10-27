package api

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func addUser(c *gin.Context) {

	var val *data.Users

	if err := c.BindJSON(&val); err != nil {
		sendResponse(2, c)
		return
	}

	createdId(val, data.DataBase)

	sendResponse(1, c)
}

func createdId(j *data.Users, db *gorm.DB) {

	if err := db.Select("Name", "Age").Create(&j).Error; err != nil {
		fmt.Println("add error", err)
		return
	}
}
