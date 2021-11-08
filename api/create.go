package api

import (
	"errors"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/Wilddogmoto/example_project/logging"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	CreateAuth struct {
		Name           string `json:"name" binding:"required,min=2,max=25"`
		Username       string `json:"username" binding:"required,min=2,max=25"`
		Password       string `json:"password" binding:"required,min=5,max=25"`
		RepeatPassword string `json:"repeatpassword" binding:"required,min=5,max=25"`
	}
)

func registrationUser(c *gin.Context) {

	var (
		logger = logging.InitLogger()
		input  CreateAuth
		out    = &data.Users{}
		user   data.Users
		err    error
		bytes  []byte
	)

	if err = c.BindJSON(&input); err != nil {
		sendResponse(2, c)
		logger.Errorf("parse json error: %v", err)
		return
	}

	if input.Password != input.RepeatPassword {
		logger.Warn("comparison password false")
		sendResponse(3, c)
		return
	}

	err = data.DataBase.First(&user, "username = ?", input.Username).Error //поиск username
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		bytes, err = bcrypt.GenerateFromPassword([]byte(input.Password), 5)
		if err != nil || bytes == nil {
			logger.Errorf("hash password error: %v", err)
			sendResponse(10, c)
			return
		}
		out.Name = input.Name
		out.Username = input.Username
		out.Password = string(bytes)

		if err = data.DataBase.Table("users").Create(&out).Error; err != nil {
			logger.Errorf("add user error: %v", err)
			sendResponse(10, c)
			return
		}
		logger.Infof("account created for username: %v,user_id: %v", out.Username, out.Id)
		sendResponse(5, c)

	case err != nil:
		logger.Errorf("search user error: %v", err)
		sendResponse(10, c)
		return

	default:
		logger.Warn("name is taken")
		sendResponse(4, c)
		return
	}
}
