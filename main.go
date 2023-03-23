package main

import (
	"github.com/4rneee/STARTHack23-sunrise-backend/controllers"
	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
    models.ConnectDatabase()

    router := gin.Default()
    router.PUT("register/", controllers.AddNewUser) 
    router.POST("login/", controllers.LoginUser)

    router.PUT("addFriendship/", controllers.AddNewFriendship)
    router.GET("getFriends/", controllers.GetActiveFriends)

    router.PUT("addPoll/", controllers.CreatePoll)
    router.PUT("putVote/", controllers.PutVote)
    router.PUT("getPolls/", controllers.GetPolls)

    router.GET("getStreams/", controllers.GetAllStreams)

    router.PUT("putComment/", controllers.PutComment)
    router.PUT("getComments/", controllers.GetComments)

    router.PUT("putClip/", controllers.PutClip)

    router.Run()
}
