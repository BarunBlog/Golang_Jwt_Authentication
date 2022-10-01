package routes

import (
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/controllers"
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

var initRouter = func() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
