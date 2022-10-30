package database

import (
	"fmt"
	"os"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	DB             *gorm.DB
	UserService    services.IUserService
	TweetService   services.ITweetService
	SessionService services.ISessionService
	ReplyService   services.IReplyService
}

func NewDBConnection() *DbConnection {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn, err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Tweet{})
	db.AutoMigrate(&entities.TweetLike{})
	db.AutoMigrate(&entities.Session{})
	db.AutoMigrate(&entities.TweetType{})
	db.AutoMigrate(&entities.TweetReply{})

	us := services.InitializedUserService(db)
	ts := services.InitializedTweetService(db)
	ss := services.InitializedSessionService(db)
	rs := services.InitializedReplyService(db)

	return &DbConnection{
		DB:             db,
		UserService:    us,
		TweetService:   ts,
		SessionService: ss,
		ReplyService:   rs,
	}
}
