package routes

import (
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/controllers"
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

var InitRouter = func() *gin.Engine {
	router := gin.Default() // Creates a new Gin router instance

	// Grouped everything under /api
	api := router.Group("/api")
	{
		api.POST("/token/", controllers.GenerateToken)
		api.POST("/user/register/", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			// We need to secure all the endpoints that will come under the api/secured/ routes.
			// We tell GIN to use the middleware that we created.
			secured.GET("/ping/", controllers.Ping)
		}
	}

	return router
}
