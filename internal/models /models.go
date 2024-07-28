package models_

// 12333434
type Users struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Books struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

type Reviews struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	BookID  uint   `json:"book_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
