package controllers

import (
	"net/http"

	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-gonic/gin"
)

type Stream struct {
    ID          uint   `json:"id" binding:"required"`
    Name        string `json:"name" binding:"required"`
    Description string `json:"description" binding:"required"`
    Thumbnail   string `json:"thumbnail" binding:"required"`
}

type CreateStream struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description" binding:"required"`
    Thumbnail   string `json:"thumbnail" binding:"required"`
}


func GetAllStreams(c *gin.Context){
    var streams []models.Stream
    if err := models.DB.Find(&streams).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, streams)
}


func AddStream(c *gin.Context){
    var newStream CreateStream

    if err := c.ShouldBindJSON(&newStream); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    err := models.DB.Create(&models.Stream{Name:newStream.Name, Description:newStream.Description, Thumbnail: newStream.Thumbnail}).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
}




