package api

import (
	"github.com/Wilddogmoto/example_project/logging"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {

	gin.SetMode(gin.ReleaseMode)
	logger := logging.InitLogger()
	router := gin.Default()

	public := router.Group("/auth")
	{
		public.POST("/reg", registrationUser)
		public.POST("/login", loginUser)
	}

	privat := router.Group("/privat", userIdentification)
	{
		privat.POST("/menu")
		privat.POST("/homepage")
	}

	if err := router.Run("localhost:8089"); err != nil {
		logger.Errorf("router run error: %v\n", err)
		return err
	}
	return nil
}
