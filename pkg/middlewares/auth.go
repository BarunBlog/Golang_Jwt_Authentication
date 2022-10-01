package middlewares

import (
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extracting the Authorization header from the HTTP context
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		err := auth.ValidaToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		// If the token is valid, the middleware allows the flow and the request reaches the required controller’s endpoint.
		context.Next()
	}
}
