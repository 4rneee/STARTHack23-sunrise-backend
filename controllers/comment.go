package controllers

import (
	"log"
	"net/http"

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

type CommentBatch struct {
	Names    []string `json:"names" binding:"required"`
	Comments []string `json:"comments" binding:"required"`
}

func PutComment(c *gin.Context) {
	var newComment Comment

	if err := c.ShouldBindJSON(&newComment); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Add comment to Buffer / DB
    err := models.DB.Create(&models.Comment{Content:newComment.Content, UserID:newComment.UserID, StreamID:newComment.StreamID}).Error
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

	//TODO: Get all comments since last pull from stream

	//TODO: Get name of user from user id
    
    var batch CommentBatch
    c.IndentedJSON(http.StatusOK, batch)
}
