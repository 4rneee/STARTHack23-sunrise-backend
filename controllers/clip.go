package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Clip struct {
	URL    string `json:"url" binding:"required"`
	UserID uint   `json:"userid" binding:"required"`
    Description string `json:"description" binding:"required"`
}

func PutClip(c *gin.Context){
    var newClip Clip

    if err := c.ShouldBindJSON(&newClip); err != nil {
        log.Println(err)
        c.Status(http.StatusBadRequest)
        return
    }

    //TODO: Save clip to db

    c.Status(http.StatusOK)
}


