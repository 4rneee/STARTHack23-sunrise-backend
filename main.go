package main

import (
	"github.com/4rneee/STARTHack23-sunrise-backend/controllers"
	"github.com/4rneee/STARTHack23-sunrise-backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
    }))

	router.PUT("register/", controllers.AddNewUser)
	router.POST("login/", controllers.LoginUser)
	router.GET("getUserData/", controllers.GetUserData)

	router.PUT("addFriendship/", controllers.AddNewFriendship)
	router.GET("getFriends/", controllers.GetActiveFriends)

	router.PUT("addPoll/", controllers.CreatePoll)
	router.PUT("putVote/", controllers.PutVote)
	router.PUT("getPolls/", controllers.GetPolls)

	router.GET("getStreams/", controllers.GetAllStreams)
	router.PUT("addStream/", controllers.AddStream)

	router.PUT("putComment/", controllers.PutComment)
	router.PUT("getComments/", controllers.GetComments)
	router.GET("getLatestMsgs", controllers.GetLatestMesseges)

	router.PUT("putClip/", controllers.PutClip)

	router.PUT("putTweet", controllers.PutTweet)
	router.GET("getTweets", controllers.GetTweet)

	router.Run()
}
