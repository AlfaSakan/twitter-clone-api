package main

import (
	"github.com/AlfaSakan/twitter-clone-api/src/database"
	"github.com/AlfaSakan/twitter-clone-api/src/handlers"
	"github.com/AlfaSakan/twitter-clone-api/src/middlewares"
	"github.com/AlfaSakan/twitter-clone-api/src/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"http://localhost:3000", "http://localhost:5173"})
	router.Use(cors.AllowAll())

	v1 := router.Group("/v1")

	conn := database.NewDBConnection()

	v1.Use(middlewares.DeserializeUser(conn.DB))

	userHandler := handlers.InitializedUserHandler(conn.UserService)
	tweetHandler := handlers.InitializedTweetHandler(conn.TweetService)
	sessionHandler := handlers.InitializedSessionHandler(conn.UserService, conn.SessionService)
	replyHandler := handlers.InitializedReplyHandler(conn.TweetService, conn.ReplyService)

	routes.User(v1, userHandler)
	routes.Tweet(v1, tweetHandler)
	routes.Session(v1, sessionHandler)
	routes.Reply(v1, replyHandler)

	router.Run(":8081")
}
