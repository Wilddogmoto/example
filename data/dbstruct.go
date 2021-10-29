package data

type (
	Users struct {
		Id       int    `gorm:"primaryKey"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	CreateAuth struct {
		Name           string `json:"name" binding:"required,min=2,max=25"`
		Username       string `json:"username" binding:"required,min=2,max=25""`
		Password       string `json:"password" binding:"required,min=5,max=25"`
		RepeatPassword string `json:"repeatpassword" binding:"required,min=5,max=25"`
	}
)
