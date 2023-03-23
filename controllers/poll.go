package controllers

import (
	"log"
	"net/http"

	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-gonic/gin"
)

type NewPoll struct {
	StreamID uint     `json:"streamid" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
}

type Vote struct {
	AnsID uint `json:"ansid" binding:"required"`
}

type RequestPoll struct {
	StreamID uint   `json:"streamid" binding:"required"`
	LastPull string `json:"lastpull" binding:"required"`
}

type Poll struct {
	PollID   uint     `json:"pollid" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
	AnsIDs   []uint   `json:"ansids" binding:"required"`
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

	//Add Poll to db
	poll := models.Poll{Question: newPoll.Question, StreamID: newPoll.StreamID}
	err := models.DB.Create(&poll).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, answer := range newPoll.Answers {
		err = models.DB.Create(models.PollAnswer{Votes: 0, Answer: answer, PollID: poll.ID}).Error
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Status(http.StatusOK)
}

func PutVote(c *gin.Context) {
	var newVote Vote

	if err := c.ShouldBindJSON(&newVote); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// Increment Answer of given Answer id of given Poll id
    var answer models.PollAnswer
    err := models.DB.First(&answer, newVote.AnsID).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

    answer.Votes += 1

    err = models.DB.Save(&answer).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func GetPolls(c *gin.Context) {
	var reqPoll RequestPoll

	if err := c.ShouldBindJSON(&reqPoll); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: Get all new or updated Polls since last pull
    var polls []models.Poll
    err := models.DB.Where("streamid = ?", reqPoll.StreamID).Find(&polls).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

    var fullPolls []Poll
    for _, poll := range polls {
        var answerModels []models.PollAnswer
        err := models.DB.Where("pollID = ?",poll.ID).Find(&answerModels).Error
        if err != nil {
            c.Status(http.StatusInternalServerError)
            return
        }

        answers := make([]string, 0)
        ansids := make([]uint, 0)
        for _, ans := range answerModels {
            answers = append(answers, ans.Answer)
            ansids = append(ansids, ans.ID)
        }
        fullPolls = append(fullPolls, Poll{PollID:poll.ID, Question:poll.Question, Answers:answers, AnsIDs: ansids})
    }


    resultPolls := StreamPolls{Polls:fullPolls}

	c.IndentedJSON(http.StatusOK, resultPolls)

}
