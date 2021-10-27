package data

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

var DataBase *gorm.DB

func FindConnect() {

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&amp",
		config.User,
		config.Password,
		config.Server,
		config.DB,
	)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println("bad connection to data base")
		panic(err)
	}
	db.AutoMigrate(&Users{}, &Hash{})

	fmt.Println("data base: connected")
	DataBase = db

}
