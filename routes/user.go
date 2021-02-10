package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-message/ent"
	"net/http"
)

// login POST {username: X, password: X} -> Create cookie
func Login(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var data login

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"OK": data})
	}
	return gin.HandlerFunc(fn)
}

// createAccount POST {username: X, password: X, profilePicture: X} -> Create account and set cookie
func Register(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type register struct {
		Username       string `json:"username" binding:"required"`
		Password       string `json:"password" binding:"required"`
		ProfilePicture string `json:"profilePicture" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var data register

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"OK": data})
	}
	return gin.HandlerFunc(fn)
}

// join POST {roomId: X} -> Check if roomID exist (must be login)
func Join(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type join struct {
		RoomId string `json:"roomID" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var data join

		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"OK": data})
	}
	return gin.HandlerFunc(fn)
}
