package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	recommendations_ "webBooks/internal/recommendations "
	repository_ "webBooks/internal/repository "
	server_ "webBooks/internal/server "
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	repo, err := repository_.NewRepository(dsn)
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	books, err := repo.GetBooks()
	if err != nil {
		log.Fatalf("failed to get books: %v", err)
	}

	rec := recommendations_.NewRecommender(books)
	srv := server_.SetupServer(repo, rec)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := srv.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
