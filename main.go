package main

import (
	"github.com/AlfaSakan/twitter-clone-api/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.DatabaseConnection()

	router.Run(":8081")
}
