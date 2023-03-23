package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
    email string
    password string
    name string
}

type returned_user struct {
    user_id string
}


func addNewUser(c *gin.Context){
    var newUser user
    
    if err:=c.BindJSON(&newUser); err != nil {
        return 
    }

    //TODO: get new user id
    return_user := returned_user{}

    //TODO: add user to data base


    c.IndentedJSON(http.StatusCreated, return_user)
}


