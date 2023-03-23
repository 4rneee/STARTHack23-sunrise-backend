package controllers

import (
	"log"
	"net/http"

	"github.com/4rneee/STARTHack23-sunrise-backend/models"
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

    err := models.DB.Create(&models.Clip{URL: newClip.URL, UserID: newClip.UserID, Description: newClip.Description})

    if err != nil {
        log.Println(err)
        c.Status(http.StatusInternalServerError)
        return
    }
    c.Status(http.StatusOK)
}


