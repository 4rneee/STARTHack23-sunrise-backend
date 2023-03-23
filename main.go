package main

import (
    "github.com/gin-gonic/gin"
    "github.com/4rneee/STARTHack23-sunrise-backend/controllers"
)

func main() {
    router := gin.Default()
    router.PUT("addUser/", controllers.AddNewUser) 
    router.PUT("addFriendship/", controllers.AddNewFriendship)
}
