package model

type UserBook struct {
	ID        int64          `json:"id" bun:"id,pk,autoincrement"`
	UserID    int64          `json:"user_id" bun:"user_id,notnull"`
	BookID    int64          `json:"book_id" bun:"book_id,notnull"`
	Status    BookStatusEnum `json:"status" bun:"status,notnull"`
	CreatedAt string         `json:"created_at" bun:"created_at,default:current_timestamp"`
}
