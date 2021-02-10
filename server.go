package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-message/config"
	"go-message/ent"
	"go-message/routes"
	"log"
)

func getDbConfig() string {
	dbConfig := config.Config.Database
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable\n",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Name,
		dbConfig.Password)
}

func main() {
	r := gin.Default()

	client, err := ent.Open("postgres", getDbConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	v1 := r.Group("/api")
	{
		// Health
		v1.GET("/health", routes.Health)

		// User management
		v1.POST("/login", routes.Login(context.Background(), client))
		v1.POST("/register", routes.Register(context.Background(), client))
		v1.POST("join", routes.Join(context.Background(), client))

		// Room management
		room := v1.Group("/room")
		{
			room.POST("/create", routes.CreateRoom(context.Background(), client))
			room.GET("/message", routes.GetRoomMessage(context.Background(), client))
			room.POST("/message", routes.PostRoomMessage(context.Background(), client))
			room.DELETE("/message", routes.DeleteRoomMessage(context.Background(), client))
			room.PUT("/message", routes.PutRoomMessage(context.Background(), client))
		}
	}

	_ = r.Run(":" + config.Config.Server.Port)
}
