package model

type User struct {
	ID        int64       `json:"id" bun:"id,pk,autoincrement"`
	Username  string      `json:"username" bun:"username,unique,notnull"`
	Email     string      `json:"email" bun:"email,unique,notnull"`
	Password  string      `json:"-" bun:"password,notnull"` // Хранить в зашифрованном виде
	Library   []*UserBook `json:"library" bun:"rel:has-many,join:id=user_id"`
	Reviews   []*Review   `json:"reviews" bun:"rel:has-many,join:id=user_id"`
	CreatedAt string      `json:"created_at" bun:"created_at,default:current_timestamp"`
}
