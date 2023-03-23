package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	Content  string `json:"content" binding:"required"`
	UserID   uint   `json:"userid" binding:"required"`
	StreamID uint   `json:"streamid" binding:"required"`
}

type RequestComment struct {
	LastPull string `json:"lastPull" binding:"required"`
	StreamID uint   `json:"streamid" binding:"required"`
}

type CommentResult struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"comment" binding:"required"`
}

func PutComment(c *gin.Context) {
	var newComment Comment

	if err := c.ShouldBindJSON(&newComment); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Add comment to Buffer / DB
	err := models.DB.Create(&models.Comment{Content: newComment.Content, UserID: newComment.UserID, StreamID: newComment.StreamID}).Error
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func GetComments(c *gin.Context) {
	var reqComment RequestComment

	if err := c.ShouldBindJSON(&reqComment); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	parsedTime, err := time.Parse(time.RFC3339Nano, reqComment.LastPull)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err)
		return
	}

	var comments []CommentResult
	err = models.DB.Raw("SELECT u.name, c.content FROM comments c, users u WHERE c.stream_id = ? AND c.user_id = u.id AND c.updated_at > ?", reqComment.StreamID, parsedTime).
		Scan(&comments).Error
    if err != nil {
        log.Println(err)
        c.Status(http.StatusInternalServerError)
        return
    }

	c.JSON(http.StatusOK, comments)
}
