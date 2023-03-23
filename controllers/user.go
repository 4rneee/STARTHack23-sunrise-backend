package controller

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

type userID struct {
    userId string
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


func addNewUser(c *gin.Context){
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

func getUserKarma(c *gin.Context){
    //TODO: get user karma and possible services
    uKarma := userKarma{}
    c.IndentedJSON(http.StatusOK, uKarma)
}

func addNewFriendship(c *gin.Context){
    var newFriendship friendship

    if err:=c.BindJSON(&newFriendship); err != nil {
        fmt.Println(err) 
        return
    }

    //TODO: save friendship

    c.Status(http.StatusOK)
}

func getActiveFriends(c *gin.Context){
    var userId userID

    if err:=c.BindJSON(&userId); err != nil {
        fmt.Println(err) 
        return
    }

    //TODO: search for active active friends 
    friends := make([]friend, 0)
    c.IndentedJSON(http.StatusOK, onlineFriends{friends:friends})
}


