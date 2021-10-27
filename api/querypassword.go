package api

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func queryPassword(c *gin.Context) {

	var val *data.Hash

	if err := c.BindJSON(&val); err != nil {
		sendResponse(2, c)
		return
	}

	passwordId(val, data.DataBase)

	sendResponse(1, c)
}

func passwordId(a *data.Hash, db *gorm.DB) {
	var result *data.Hash

	if err := db.Table("hashes", &data.Hash{}).Find(&a).Take(&result).Error; err != nil {
		fmt.Println("query error", err)
		return
	}
	fmt.Println(result.Password)
}
