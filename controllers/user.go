package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type loginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type userID struct {
	UserId uint `json:"userid" binding:"required"`
}

type userKarma struct {
	Points  uint   `json:"points" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type friendship struct {
	User1 uint `json:"user1" binding:"required"`
	User2 uint `json:"user2" binding:"required"`
}

type friend struct {
	Name   string `json:"name" binding:"required"`
	ViewId uint   `json:"viewid" binding:"required"`
}

type onlineFriends struct {
	Friends []friend `json:"friends" binding:"required"`
}

func AddNewUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}

	//TODO: get new user id
	user := userID{}

	//TODO: add user to data base

	c.IndentedJSON(http.StatusCreated, user)
}

func LoginUser(c *gin.Context) {
	var newUser loginUser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: Check if user credentials are stored and get userId
	//check
	userId := userID{UserId: 0}

	c.IndentedJSON(http.StatusOK, userId)
}

func GetUserKarma(c *gin.Context) {
	//TODO: get user karma and possible services
	uKarma := userKarma{}
	c.IndentedJSON(http.StatusOK, uKarma)
}

func AddNewFriendship(c *gin.Context) {
	var newFriendship friendship

	if err := c.ShouldBindJSON(&newFriendship); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: save friendship

	c.Status(http.StatusOK)
}

func GetActiveFriends(c *gin.Context) {
	var userId userID

	if err := c.ShouldBindJSON(&userId); err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	//TODO: search for active active friends
	friends := make([]friend, 0)
	c.IndentedJSON(http.StatusOK, onlineFriends{Friends: friends})
}
