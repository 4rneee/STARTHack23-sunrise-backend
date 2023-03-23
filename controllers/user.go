package controllers

import (
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"strings"

	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	Name     string `json:"name" binding:"required"`
	StreamID uint   `json:"streamid" binding:"required"`
}

type onlineFriends struct {
	Friends []friend `json:"friends" binding:"required"`
}

func AddNewUser(c *gin.Context) {
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	// check if email is valid
	if email, err := mail.ParseAddress(input.Email); err == nil {
		input.Email = email.Address
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inDB models.User
	err := models.DB.
		Table("users").
		Where("email = ?", input.Email).
		First(&inDB).
		Error

	// no error => a user with the email exists
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "an account with this email already exists"})
		return
	}

	if input.Name = strings.TrimSpace(input.Name); input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name must not be empty"})
		return
	}

	// TODO: proper password check
	if input.Password = strings.TrimSpace(input.Password); input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must not be empty"})
		return
	} else if len(input.Password) > 72 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password can not be longer than 72 bytes"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insert := models.User{
		ID:       0, // will be automatically set by DB
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
		Points:   0,
		Friends:  nil,
		StreamID: 0,
	}

	err = models.DB.
		Create(&insert).
		Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func LoginUser(c *gin.Context) {
	var input loginUser

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	var inDB models.User
	err := models.DB.First(&inDB, "email = ?", input.Email).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		log.Println(err.Error())
		return
	}

	if err = bcrypt.CompareHashAndPassword(inDB.Password, []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": inDB.ID})
}

func GetUserKarma(c *gin.Context) {
	//TODO: get user karma and possible services
	uKarma := userKarma{}
	c.IndentedJSON(http.StatusOK, uKarma)
}

func AddNewFriendship(c *gin.Context) {
	var newFriendship friendship

	if err := c.ShouldBindJSON(&newFriendship); err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	var u1 models.User
	if err := models.DB.First(&u1, newFriendship.User1).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user1 does not exist"})
		return
	}

	var u2 models.User
	if err := models.DB.First(&u2, newFriendship.User2).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user2 does not exist"})
		return
	}

	if err := models.DB.Exec("INSERT INTO \"friends\" VALUES (?, ?);", newFriendship.User1, newFriendship.User2).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "user2 is already a friend of user1"})
		return
	}

	c.Status(http.StatusOK)
}

func GetActiveFriends(c *gin.Context) {
	id_s, ok := c.GetQuery("id")

	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	var friends []models.User
	err = models.DB.Raw("SELECT fs.* FROM users u JOIN friends f ON u.id=f.user_id JOIN users fs ON fs.id=f.friend_id WHERE u.id = ? AND fs.stream_id > 0", id).Scan(&friends).Error

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	friends_ret := make([]friend, 0, len(friends))

	for _, f := range friends {
		friends_ret = append(friends_ret, friend{Name: f.Name, StreamID: f.StreamID})
	}

    c.JSON(http.StatusOK, friends_ret)
}
