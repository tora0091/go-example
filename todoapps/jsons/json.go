package jsons

import "time"

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_ad"`
	UpdatedAt time.Time `json:"updated_ad"`
}

type Todos []Todo
