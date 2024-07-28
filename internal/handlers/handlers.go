package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	models_ "webBooks/internal/models "
	recommendations_ "webBooks/internal/recommendations "
	repository_ "webBooks/internal/repository "
)

type Handler struct {
	repo *repository_.Repository
	rec  *recommendations_.Recommender
}

func NewHandler(repo *repository_.Repository, rec *recommendations_.Recommender) *Handler {
	return &Handler{repo: repo, rec: rec}
}

func (h *Handler) Register(c *gin.Context) {
	var user models_.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered succeddfully"})

}

func (h *Handler) Login(c *gin.Context) {
	var user models_.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := h.repo.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if dbUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid operation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h *Handler) AddBook(c *gin.Context) {
	var book models_.Books
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.repo.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) AddReview(c *gin.Context) {
	var review models_.Reviews
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "review added seccessfully"})
}

func (h *Handler) GetReviews(c *gin.Context) {
	bookIDStr := c.Query("book_id")
	if bookIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book id is required"})
		return
	}

	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	reviews, err := h.repo.GetReviewsByBookID(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *Handler) RecommendBooks(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	prefs := map[string]float64{
		"fiction": 1.0,
		"fantasy": 2.0,
		"romance": 0.5,
	}

	recommendedBooks := h.rec.Recommend(prefs)
	c.JSON(http.StatusOK, recommendedBooks)
}
