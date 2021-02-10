package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-message/ent"
	"net/http"
)

// createRoom POST -> Create room (must be login)
func CreateRoom(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	}
	return gin.HandlerFunc(fn)
}

// roomMessage GET {roomId: X} -> Return list of message from roomID
func GetRoomMessage(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type roomMessage struct {
		RoomID string `form:"roomId" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var data roomMessage

		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"OK": data})
	}
	return gin.HandlerFunc(fn)
}

// roomMessage POST {roomId : X} -> Add a message to roomID
func PostRoomMessage(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type roomMessageQuery struct {
		RoomID string `form:"roomId" binding:"required"`
	}

	type roomMessageBody struct {
		Message string `form:"message" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var dataQuery roomMessageQuery
		var dataBody roomMessageBody

		if err := c.ShouldBindQuery(&dataQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&dataBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"OK": dataBody})
	}
	return gin.HandlerFunc(fn)
}

// roomMessage Delete {roomId: X, messageId: X} -> Delete a message from roomID (check if sender has the good ID)
func DeleteRoomMessage(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type roomMessage struct {
		RoomID    string `form:"roomId" binding:"required"`
		MessageID string `form:"messageId" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var data roomMessage

		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"OK": data})
	}
	return gin.HandlerFunc(fn)
}

// roomMessage PUT {room:Id: X, messageId: X} -> Update a message from roomID
func PutRoomMessage(ctx context.Context, client *ent.Client) gin.HandlerFunc {
	type roomMessageQuery struct {
		RoomID    string `form:"roomId" binding:"required"`
		MessageID string `form:"messageId" binding:"required"`
	}

	type roomMessageBody struct {
		Message string `form:"message" binding:"required"`
	}

	fn := func(c *gin.Context) {
		var dataQuery roomMessageQuery
		var dataBody roomMessageBody

		if err := c.ShouldBindQuery(&dataQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&dataBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dataBody})
	}
	return gin.HandlerFunc(fn)
}
