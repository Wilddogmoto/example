package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	router := gin.Default()

	public := router.Group("/auth")
	{
		public.POST("/reg", regUser)

		public.POST("/login", loginUser)

	}

	privat := router.Group("/privat")
	{
		privat.POST("/menu", userIdentification)
	}

	if err := router.Run("localhost:8089"); err != nil {
		fmt.Println("router run error")
		return err
	}

	return nil
}
