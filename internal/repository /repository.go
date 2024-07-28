package repository_

import (
	"database/sql"
	"fmt"
	models_ "webBooks/internal/models "
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dataSourceName string) (*Repository, error) {
	db, err := sql.Open("webBook", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) CreateUser(user *models_.Users) error {
	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	_, err := r.db.Exec(query, user.Username, user.Password)
	return err
}

func (r *Repository) GetUserByUsername(username string) (*models_.Users, error) {
	query := `SELECT id, username, password FROM users where username = ?`
	row := r.db.QueryRow(query, username)

	var user models_.Users
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) AddBook(book *models_.Books) error {
	query := `INSERT INTO books (title, author, genre, description) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, book.Title, book.Author, book.Genre, book.Description)
	return err
}

func (r *Repository) GetBooks() ([]*models_.Books, error) {
	query := `SELECT id, title, author, genre, description FROM books`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models_.Books
	for rows.Next() {
		var book models_.Books
		if err := rows.Scan(&book.ID, &book.Title, &book.Genre, &book.Author, &book.Description); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *Repository) CreateReview(review *models_.Reviews) error {
	query := `INSERT INTO reviews (user_id, book_id, rating, comment) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, review.UserID, review.BookID, review.Rating, review.Comment)
	if err != nil {
		return err
	}

	return r.updateBookRating(review.BookID)
}

func (r *Repository) GetReviewsByBookID(bookID uint) ([]*models_.Reviews, error) {
	query := `SELECT id, user_id, book_id, rating, comments FROM reviews where book_id = ?`
	rows, err := r.db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*models_.Reviews
	for rows.Next() {
		var review models_.Reviews
		if err := rows.Scan(&review.ID, &review.UserID, &review.BookID, &review.Rating, &review.Comment); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (r *Repository) updateBookRating(bookID uint) error {
	query := `SELECT AVG(rating) FROM reviews WHERE book_id = ?`
	row := r.db.QueryRow(query, bookID)

	var avgRating float64
	if err := row.Scan(&avgRating); err != nil {
		return err
	}

	updateQuery := `UPDATE books SET rating = ? WHERE id = ?`
	_, err := r.db.Exec(updateQuery, avgRating, bookID)
	return err
}

func (r *Repository) CreateBook(book *models_.Books) error {
	query := `INSERT INTO books (title, author, genre, description, rating) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, book.Title, book.Author, book.Genre, book.Description, book.Rating)
	return err
}
