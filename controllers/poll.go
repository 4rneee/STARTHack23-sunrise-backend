package controllers

import (
	"log"
	"net/http"
	"time"

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

type Answer struct {
	AnsID  uint   `json:"id"`
	answer string `json:"choice"`
	Votes  uint   `json:"votes"`
}
type Poll struct {
	PollID   uint     `json:"pollid" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []Answer `json:"answers"`
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
		log.Println(err)
		return
	}

	for _, answer := range newPoll.Answers {
		err = models.DB.Create(&models.PollAnswer{Votes: 0, Answer: answer, PollID: poll.ID}).Error
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
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
		log.Println(err)
		return
	}

	answer.Votes += 1

	err = models.DB.Save(&answer).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
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
	parsedTime, err := time.Parse(time.RFC3339, reqPoll.LastPull)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Get all new or updated Polls since last pull
	var polls []models.Poll
	err = models.DB.Where("stream_iD = ?", reqPoll.StreamID).Where("updated_at > ?", parsedTime).Find(&polls).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var fullPolls []Poll = make([]Poll, 0, len(polls))
	for _, poll := range polls {
		var answerModels []models.PollAnswer
		err := models.DB.Where("poll_id = ?", poll.ID).Find(&answerModels).Error
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
        log.Println(answerModels)

		answers := make([]Answer, 0, len(answerModels))
		for _, ans := range answerModels {
			answers = append(answers, Answer{AnsID: ans.ID, answer: ans.Answer, Votes: ans.Votes})
		}
		fullPolls = append(fullPolls, Poll{PollID: poll.ID, Question: poll.Question, Answers: answers})
	}


	c.IndentedJSON(http.StatusOK, fullPolls)

}
