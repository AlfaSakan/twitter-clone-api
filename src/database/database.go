package database

import (
	"fmt"
	"os"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/AlfaSakan/twitter-clone-api/src/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	UserService services.IUserService
}

func NewDBConnection() *DbConnection {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase, dbPort)
	// dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn, err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.User{})

	ur := repositories.InitializedUserRepository(db)
	us := services.InitializedUserService(ur)

	return &DbConnection{
		UserService: us,
	}
}
