package db

import (
	"fmt"
	"gorm.io/gorm"
)

func deleteId(db *gorm.DB) {

	db.Where("name = ?", "Jin").Delete(&Users{})
	fmt.Println("user:deleted")
}
