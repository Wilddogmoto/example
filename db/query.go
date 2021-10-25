package db

import (
	"gorm.io/gorm"
)

func queryId(db *gorm.DB) {

	db.Where(&Users{Name: "Mac", Age: 25}).First(&Users{})

}
