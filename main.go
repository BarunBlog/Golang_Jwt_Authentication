package main

import (
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/routes"
)

func main() {
	router := routes.InitRouter()
	router.Run("127.0.0.1:4000")
}
