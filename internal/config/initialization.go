package config

import (
	"fmt"
	"log"
	"os"
	"stocksync_api/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeEverything() *gin.Engine {
	db := initializeDB()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	router := gin.Default()
	v1 := router.Group("api/v1")

	user.RegisterRoutes(v1, *userHandler)
	return router
}

func initializeDB() *gorm.DB {
	env := getEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.Host,
		env.User,
		env.Password,
		env.Name,
		env.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database is failed to connect.")
	}

	return db
}

func getEnv() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting .env file")
	}

	return Env{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}
