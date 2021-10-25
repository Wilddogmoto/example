package db

type (
	Users struct {
		Id   uint16 `gorm:"primaryKey"`
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}
)

var user = Users{
	Name: "Mac",
	Age:  25,
}
