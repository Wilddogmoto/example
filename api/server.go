package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	public := router.Group("/public")
	{
		public.POST("/add_user", addUser)

		public.POST("/find_user", queryUser)

		public.POST("/delete_user", deleteUser)

	}

	password := router.Group("/password")
	{
		password.POST("/create", addPassword)

		password.POST("/find", queryPassword)
	}

	router.Run("localhost:8089")
}
