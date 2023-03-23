package main

import (
    "github.com/gin-gonic/gin"
    "github.com/4rneee/STARTHack23-sunrise-backend/controllers"
)

func main() {
    router := gin.Default()
    router.PUT("addUser/", controllers.AddNewUser) 
    router.GET("login/", controllers.LoginUser)

    router.PUT("addFriendship/", controllers.AddNewFriendship)
    router.GET("getFriends/", controllers.GetActiveFriends)

    router.Run("localhost:8080")
}
