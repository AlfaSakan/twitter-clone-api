package main

import (
	"github.com/AlfaSakan/twitter-clone-api/src/database"
	"github.com/AlfaSakan/twitter-clone-api/src/handlers"
	"github.com/AlfaSakan/twitter-clone-api/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	conn := database.NewDBConnection()

	userHandler := handlers.InitializedUserHandler(conn.UserService)
	routes.User(v1, userHandler)

	router.Run(":8081")
}
