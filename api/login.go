package api

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type (
	TokenClaims struct {
		UserId int
		jwt.StandardClaims
	}
)

var (
	inputLogin data.Users

	account = &data.Users{}
)

func loginUser(c *gin.Context) {

	if err := c.BindJSON(&inputLogin); err != nil {

		sendResponse(2, c)
		fmt.Println(err)
		return
	}

	if queryUser(inputLogin.Username, inputLogin.Password, c, data.DataBase) == true {

		createJwt(c)
	}

}

func queryUser(username, password string, c *gin.Context, db *gorm.DB) bool { // проверка на существование учетной записи

	if err := db.Table("users").Where("username = ?", username).First(account).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			sendResponse(6, c)
			return false
		}

		fmt.Println(err)
		return false

	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil { // получаем хэш пароля

		if err == bcrypt.ErrMismatchedHashAndPassword {
			sendResponse(7, c)
			return false
		}

		fmt.Println(err)
		return false

	}

	return true
}

func createJwt(c *gin.Context) {

	claims := &TokenClaims{UserId: account.Id}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signaturekey))
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": ss,
		"id":    account.Id,
	})

}
