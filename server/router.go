package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ltadrian/go-gin-dynamodb-boilerplate/controllers"
	"github.com/ltadrian/go-gin-dynamodb-boilerplate/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("/v1")
	{
		userGroup := v1.Group("user")
		// set jwt middleware on user group route
		userGroup.Use(middlewares.AuthMiddleware())
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.FindUserByID)
			userGroup.POST("/", user.AddUser)
		}
		tokenGroup := v1.Group("token")
		{
			token := new(controllers.TokenController)
			tokenGroup.GET("/", token.GetToken)
		}
	}

	return router
}
