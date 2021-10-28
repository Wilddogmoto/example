package api

import (
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var inputLogin data.Users

func loginUser(c *gin.Context) {

	if err := c.BindJSON(&input); err != nil {
		sendResponse(2, c)
		return
	}
	queryUser(inputLogin.Username, inputLogin.Password, c, data.DataBase)
}

func queryUser(username, password string, c *gin.Context, db *gorm.DB) {

	account := &data.Users{}

	if err := db.Table("users").Where("username = ?", username).First(account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			sendResponse(6, c)
		}
		sendResponse(2, c)
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		sendResponse(7, c)
	}

}
