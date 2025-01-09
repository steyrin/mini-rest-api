package model

type Book struct {
	ID          int64       `json:"id" bun:"id,pk,autoincrement"`
	Name        string      `json:"name" bun:"name,notnull"`
	Genre       string      `json:"genre" bun:"genre,notnull"`
	Description string      `json:"description" bun:"description"`
	Year        int64       `json:"year" bun:"year,notnull"`
	Rating      float64     `json:"rating" bun:"rating,default:0"`
	Price       float64     `json:"price" bun:"price,notnull"`
	CreatedAt   string      `json:"created_at" bun:"created_at,default:current_timestamp"`
	UserBooks   []*UserBook `json:"user_books" bun:"rel:has-many,join:id=book_id"`
	Reviews     []*Review   `json:"reviews" bun:"rel:has-many,join:id=book_id"`
}
