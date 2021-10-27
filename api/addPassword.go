package api

import (
	"errors"
	"fmt"
	"github.com/Wilddogmoto/example_project/data"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

var ErrHashTooShort = errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password")
var ErrMismatchedHashAndPassword = errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password")

func addPassword(c *gin.Context) {

	var val *data.Hash
	if err := c.BindJSON(&val); err != nil {
		sendResponse(2, c)
		return
	}

	HashPassword(val.Password)

	sendResponse(1, c)

}

func HashPassword(password string) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil || bytes == nil {
		fmt.Println("hash password error", err)
	}
	b := data.Hash{Password: string(bytes)}

	creatPass(b, data.DataBase)

}

func creatPass(c data.Hash, db *gorm.DB) {

	if err := db.Select("Id", "Password").Create(&c).Error; err != nil {
		fmt.Println("add error", err)
		return
	}
}

/*func CreateHash(key string) string{
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err !=nil{
		fmt.Println("hash password error",err)
	}
	return string(bytes), err
}*/
