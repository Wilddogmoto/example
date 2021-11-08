package api

import (
	"errors"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/Wilddogmoto/example_project/logging"
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

const (
	signaturekey = "full" // ключ зашифровки \ расшифровки
)

func loginUser(c *gin.Context) {

	var (
		logger     = logging.InitLogger()
		account    = &data.Users{}
		inputLogin data.Users
		err        error
		outToken   string
	)

	if err = c.BindJSON(&inputLogin); err != nil {
		sendResponse(2, c)
		logger.Errorf("parse json error: %v", err)
		return
	}

	if err = data.DataBase.Where("username = ?", inputLogin.Username).First(account).Error; err != nil { // проверка на существование учетной записи
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warnf("invalid username: %v", err)
			sendResponse(6, c)
			return
		}
		logger.Errorf("search user error: %v", err)
		sendResponse(10, c)
		return

	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(inputLogin.Password)); err != nil { // проверка пароля
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			logger.Warnf("wrong password: %v", err)
			sendResponse(7, c)
			return
		}

		logger.Errorf("hash password error: %v", err)
		sendResponse(10, c)
		return

	}

	claims := &TokenClaims{UserId: account.Id} //создание jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	outToken, err = token.SignedString([]byte(signaturekey))
	if err != nil {
		logger.Errorf("token create error: %v", err)
		sendResponse(10, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"Token":   outToken,
		"id":      account.Id,
	})

	logger.Infof("loginUser: login, token gived for userId - %v", account.Id)
}
