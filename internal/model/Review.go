package model

type Review struct {
	ID        int64  `json:"id" bun:"id,pk,autoincrement"`
	UserID    int64  `json:"user_id" bun:"user_id,notnull"` // ID пользователя, оставившего отзыв
	BookID    int64  `json:"book_id" bun:"book_id,notnull"` // ID книги, к которой написан отзыв
	Rating    int    `json:"rating" bun:"rating,notnull"`   // Оценка книги (например, от 1 до 5)
	Comment   string `json:"comment" bun:"comment"`         // Текст отзыва
	CreatedAt string `json:"created_at" bun:"created_at,default:current_timestamp"`
}
