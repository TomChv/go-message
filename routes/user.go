package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// login POST {username: X, password: X} -> Create cookie
func Login(c *gin.Context) {
	type login struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var data login

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}

// createAccount POST {username: X, password: X, profilePicture: X} -> Create account and set cookie
func Register(c *gin.Context) {
	type register struct {
		Username       string `json:"username" binding:"required"`
		Password       string `json:"password" binding:"required"`
		ProfilePicture string `json:"profilePicture" binding:"required"`
	}

	var data register

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}

// join POST {roomId: X} -> Check if roomID exist (must be login)
func Join(c *gin.Context) {
	type join struct {
		RoomId string `json:"roomId" binding:"required"`
	}

	var data join

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}
