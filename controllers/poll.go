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

type Poll struct {
	PollID   uint     `json:"pollid" binding:"required"`
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
	AnsIDs   []uint   `json:"ansids" binding:"required"`
	Votes    []uint   `json:"votes" binding:"required"`
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
    parsedTime, err := time.Parse(time.RFC3339Nano, reqPoll.LastPull)
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

	var fullPolls []Poll
	for _, poll := range polls {
		var answerModels []models.PollAnswer
		err := models.DB.Where("poll_iD = ?", poll.ID).Find(&answerModels).Error
		if err != nil {
			c.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		answers := make([]string, 0)
		ansids := make([]uint, 0)
		votes := make([]uint, 0)
		for _, ans := range answerModels {
			answers = append(answers, ans.Answer)
			ansids = append(ansids, ans.ID)
			votes = append(votes, ans.Votes)
		}
        fullPolls = append(fullPolls, Poll{PollID: poll.ID, Question: poll.Question, Answers: answers, AnsIDs: ansids, Votes:votes})
	}

	resultPolls := StreamPolls{Polls: fullPolls}

	c.IndentedJSON(http.StatusOK, resultPolls)

}
