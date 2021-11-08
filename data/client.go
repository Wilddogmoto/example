package data

import (
	"fmt"
	"github.com/Wilddogmoto/example_project/logging"
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

var (
	config = Account{
		//c:= "wild:777@tcp(127.0.0.1:3306)/mytestdb?charset=utf8mb4&parseTime=True&loc=Local"
		Server:   "127.0.0.1:3306",
		User:     "wild",
		Password: "777",
		DB:       "mytestdb",
	}

	DataBase *gorm.DB
	err      error
	logger   = logging.InitLogger()
)

func DBConnect() error {

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&amp",
		config.User,
		config.Password,
		config.Server,
		config.DB,
	)

	DataBase, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		logger.Error("bad connection to database")
		return err
	}

	logger.Info("DBConnect func: Database connected!")
	return nil
}
