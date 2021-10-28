package test

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/api"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func deleteUser(c *gin.Context) {

	var val *data.Users

	if err := c.BindJSON(&val); err != nil {
		api.sendResponse(2, c)
		return
	}

	deleteId(val, data.DataBase)

	api.sendResponse(1, c)
}

func deleteId(j *data.Users, db *gorm.DB) {

	if err := db.Delete(&data.Users{}, j).Error; err != nil {
		fmt.Println("delete error", err)
		return
	}

}
