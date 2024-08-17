package todos

import "time"

type Todo struct {
	ID        int       `json:"todo_id" bun:"todo_id,pk,autoincrement,notnull"`
	Title     string    `json:"title" bun:",notnull,type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp,type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp,type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" bun:",type:datetime,default:null"`
}
