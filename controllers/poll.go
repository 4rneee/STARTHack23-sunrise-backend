package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Poll struct {
	ID       uint     `json:"id" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
}

type Vote struct {
	PollID uint `json:"pollid" binding:"required"`
	AnsID  uint `json:"ansid" binding:"required"`
}

func CreatePoll(c *gin.Context){
    var newPoll Poll

    if err := c.ShouldBindJSON(&newPoll); err != nil{
        log.Println(err)
        c.Status(http.StatusBadRequest)
        return
    }

    //TODO: Add Poll to db

    c.Status(http.StatusOK)
}

func PutVote(c *gin.Context){
    var newVote Poll

    if err := c.ShouldBindJSON(&newVote); err != nil{
        log.Println(err)
        c.Status(http.StatusBadRequest)
        return
    }

    //TODO: Increment Answer of given Answer id of given Poll id

    c.Status(http.StatusOK)
}
