package api

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	input data.CreateAuth
	out   data.Users
)

func regUser(c *gin.Context) {

	if err := c.BindJSON(&input); err != nil {
		sendResponse(2, c)
		return
	}

	if input.Password != input.RepeatPassword {
		sendResponse(3, c)
		return
	}

	if searchUser(input.Username, data.DataBase, c) == true {
		hashPassword(input.Password)
		createdId(out, data.DataBase)
		sendResponse(5, c)
	}

}
func searchUser(a string, db *gorm.DB, c *gin.Context) bool {

	if err := db.Table("users").Where("username = ?", a).First(&data.Users{}).Error; err != nil {
		fmt.Println(err)
		return true
	}
	sendResponse(4, c)
	return false
}

func hashPassword(password string) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil || bytes == nil {
		fmt.Println("hash password error", err)
	}

	out.Name = input.Name
	out.Username = input.Username
	out.Password = string(bytes)
}

func createdId(val data.Users, db *gorm.DB) {

	if err := db.Table("users").Create(&val).Error; err != nil {
		fmt.Println("add error", err)
		return
	}
}
