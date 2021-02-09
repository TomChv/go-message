package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// createRoom POST -> Create room (must be login)
func CreateRoom(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// roomMessage GET {roomId: X} -> Return list of message from roomID
func GetRoomMessage(c *gin.Context) {
	type roomMessage struct {
		RoomID string `form:"roomId" binding:"required"`
	}

	var data roomMessage
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}

// roomMessage POST {roomId : X} -> Add a message to roomID
func PostRoomMessage(c *gin.Context) {
	type roomMessage struct {
		RoomID  string `form:"roomId" binding:"required"`
		Message string `form:"message" binding:"required"`
	}

	var data roomMessage
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}

// roomMessage Delete {roomId: X, messageId: X} -> Delete a message from roomID (check if sender has the good ID)
func DeleteRoomMessage(c *gin.Context) {
	type roomMessage struct {
		RoomID    string `form:"roomId" binding:"required"`
		MessageID string `form:"messageId" binding:"required"`
	}

	var data roomMessage
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}

// roomMessage PUT {room:Id: X, messageId: X} -> Update a message from roomID
func PutRoomMessage(c *gin.Context) {
	type roomMessage struct {
		RoomID    string `form:"roomId" binding:"required"`
		MessageID string `form:"messageId" binding:"required"`
	}

	var data roomMessage
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"OK": data})
}
