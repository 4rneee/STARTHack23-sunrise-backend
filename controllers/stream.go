package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Stream struct {
    ID          uint   `json:"id" binding:"required"`
    Name        string `json:"name" binding:"required"`
    Description string `json:"description" binding:"required"`
    Thumbnail   string `json:"thumbnail" binding:"required"`
}


func GetAllStreams(c *gin.Context){
    //TODO: get all streams 
    var streams []Stream

    c.IndentedJSON(http.StatusOK, streams)
}




