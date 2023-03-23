package controllers

import (
	"log"
	"net/http"

	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-gonic/gin"
)

type tweetInput struct {
	UserName string `json:"username" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

func PutTweet(c *gin.Context) {
    var input tweetInput
    if err := c.ShouldBindJSON(&input); err != nil {
        log.Println(err)
        c.Status(http.StatusBadRequest)
        return
    }
    if err := models.DB.Create(&models.Tweet{UserName: input.UserName, Content: input.Content}).Error; err != nil {
        log.Println(err)
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusOK)
}

func GetTweet(c *gin.Context) {
	tag, ok := c.GetQuery("hashtag")

	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

    var tweets []models.Tweet
    // please don't abuse SQL injection
    if err := models.DB.Where("content LIKE '%#" + tag + "%'").Find(&tweets).Error; err != nil {
        log.Println(err)
        c.Status(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusOK, tweets)
}
