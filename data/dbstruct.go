package data

type (
	Users struct {
		Id       int    `gorm:"primaryKey"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
