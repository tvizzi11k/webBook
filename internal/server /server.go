package server_

import (
	"github.com/gin-gonic/gin"
	"webBooks/internal/handlers"
	recommendations_ "webBooks/internal/recommendations "
	repository_ "webBooks/internal/repository "
)

func SetupServer(repo *repository_.Repository, rec *recommendations_.Recommender) *gin.Engine {
	handler := handlers.NewHandler(repo, rec)
	router := gin.Default()

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.POST("/books", handler.AddBook)
	router.GET("/books", handler.GetBooks)
	router.POST("/reviews", handler.AddReview)
	router.GET("/reviews", handler.GetReviews)
	router.GET("/recommend/:user_id", handler.RecommendBooks)

	return router
}
