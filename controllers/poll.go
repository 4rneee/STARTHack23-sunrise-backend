package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewPoll struct {
	StreamID uint     `json:"streamid" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
}

type Vote struct {
	StreamID uint `json:"streamid" binding:"required"`
	PollID   uint `json:"pollid" binding:"required"`
	AnsID    uint `json:"ansid" binding:"required"`
}

type RequestPoll struct {
	StreamID uint `json:"streamid" binding:"required"`
    LastPull string `json:"lastpull" binding:"required"` 
}

type Poll struct {
	PollID   uint `json:"pollid" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
}

type StreamPolls struct {
    Polls []Poll `json:"polls" binding:"required"`
}

func CreatePoll(c *gin.Context) {
	var newPoll NewPoll

	if err := c.ShouldBindJSON(&newPoll); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: Add Poll to db

	c.Status(http.StatusOK)
}

func PutVote(c *gin.Context) {
	var newVote NewPoll

	if err := c.ShouldBindJSON(&newVote); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: Increment Answer of given Answer id of given Poll id

	c.Status(http.StatusOK)
}

func GetPolls( c *gin.Context){
    var reqPoll RequestPoll

    if err := c.ShouldBindJSON(&reqPoll); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

    //TODO: Get all new or updated Polls since last pull

    var polls StreamPolls

    c.IndentedJSON(http.StatusOK, polls)
 
}
