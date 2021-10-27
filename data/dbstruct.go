package data

type (
	Users struct {
		Id   uint16 `gorm:"primaryKey"`
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}
)

type Hash struct {
	Id       uint16 `gorm:"primaryKey"`
	Password string `json:"password"`
}

/*var user = Users{
	Name: "Jack",
	Age:  26,
}*/
