package controllers

import (
	"net/http"
    "fmt"
	"github.com/gin-gonic/gin"
)

type user struct {
    email string
    password string
    name string
}

type loginUser struct {
    email string
    password string
}

type userID struct {
    userId uint
}

type userKarma struct {
    points uint
    message string
}

type friendship struct {
    user1 uint
    user2 uint
}

type friend struct {
    name string
    viewId uint
}

type onlineFriends struct{
    friends []friend
}


func AddNewUser(c *gin.Context){
    var newUser user
    
    if err:=c.BindJSON(&newUser); err != nil {
        fmt.Println(err) 
        return 
    }

    //TODO: get new user id
    user := userID{}

    //TODO: add user to data base


    c.IndentedJSON(http.StatusCreated, user)
}

func LoginUser(c *gin.Context){
    var newUser loginUser

    if err:=c.BindJSON(&newUser); err != nil {
        fmt.Println(err) 
        c.Status(http.StatusBadRequest)
        return
    }

    //TODO: Check if user credentials are stored and get userId
    //check
    userId := userID{userId:0}

    c.IndentedJSON(http.StatusOK, userId)  
}


func GetUserKarma(c *gin.Context){
    //TODO: get user karma and possible services
    uKarma := userKarma{}
    c.IndentedJSON(http.StatusOK, uKarma)
}

func AddNewFriendship(c *gin.Context){
    var newFriendship friendship

    if err:=c.BindJSON(&newFriendship); err != nil {
        fmt.Println(err) 
        c.Status(http.StatusBadRequest)
        return
    }

    //TODO: save friendship

    c.Status(http.StatusOK)
}

func GetActiveFriends(c *gin.Context){
    var userId userID

    if err:=c.BindJSON(&userId); err != nil {
        fmt.Println(err) 
        c.Status(http.StatusBadRequest)
        return
    }

    //TODO: search for active active friends 
    friends := make([]friend, 0)
    c.IndentedJSON(http.StatusOK, onlineFriends{friends:friends})
}


