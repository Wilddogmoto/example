package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	Account struct {
		Server   string
		User     string
		Password string
		DB       string
	}
)

var config = Account{
	//c:= "wild:777@tcp(127.0.0.1:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
	Server:   "127.0.0.1:3306",
	User:     "wild",
	Password: "777",
	DB:       "mytestdb",
}

func Connect() {

	getConfig(config)

}

func getConfig(config Account) {

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&amp",
		config.User,
		config.Password,
		config.Server,
		config.DB,
	)
	findeConnect(connection)
}

func findeConnect(con string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB: Connected")

	return db

}
