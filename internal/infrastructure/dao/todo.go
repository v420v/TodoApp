package dao

import "time"

type Todo struct {
	Id        string    `bun:"todo_id,pk,autoincrement,notnull"`
	Title     string    `bun:",notnull,type:varchar(255)"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp,type:timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp,type:timestamp"`
	DeletedAt time.Time `bun:",type:datetime,default:null"`
}
