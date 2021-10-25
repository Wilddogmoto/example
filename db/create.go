package db

import (
	"fmt"
	"gorm.io/gorm"
)

func createdId(db *gorm.DB) {

	db.Select("Name", "Age").Create(&user)
	fmt.Println("user: added")

}
