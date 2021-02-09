package main

import (
	"github.com/gin-gonic/gin"
	"go-message/routes"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		// Health
		v1.GET("/health", routes.Health)

		// User management
		v1.POST("/login", routes.Login)
		v1.POST("/register", routes.Register)
		v1.POST("join", routes.Join)

		// Room management
		room := v1.Group("/room")
		{
			room.POST("/create", routes.CreateRoom)
			room.GET("/message", routes.GetRoomMessage)
			room.POST("/message", routes.PostRoomMessage)
			room.DELETE("/message", routes.DeleteRoomMessage)
			room.PUT("/message", routes.PutRoomMessage)
		}
	}

	_ = r.Run()
}
