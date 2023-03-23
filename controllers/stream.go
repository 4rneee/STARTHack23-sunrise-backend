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


func GetAllStreams(c *gin.Context){
    var streams []models.Stream
    models.DB.Find(&streams)

    c.JSON(http.StatusOK, streams)
}




