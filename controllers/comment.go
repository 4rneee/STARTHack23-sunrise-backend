package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Comment struct {
    Content  string `json:"content" binding:"required"`
    UserID   uint   `json:"userid" binding:"required"`
    StreamID uint `json:"streamid" binding:"required"`
}

func PutComment(c *gin.Context){
    var newComment Comment

    if err := c.ShouldBindJSON(&newComment); err != nil {
        log.Println(err)
        c.Status(http.StatusBadRequest)
        return
    }

    // TODO: Add comment to Buffer / DB

     c.Status(http.StatusOK)
}

