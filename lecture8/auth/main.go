package main

import (
	"github.com/gin-gonic/gin"
	"justCode/lecture8/auth/api"
	"justCode/lecture8/auth/middleware"
)

func main() {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware)

	group := r.Group("/api")

	group.POST("/endpoint", api.CreateUser)

	group.GET("/endpoint/:id", api.OutUser)

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
