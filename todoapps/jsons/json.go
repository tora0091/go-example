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

type JSONStatusOKResponse struct {
	Status int `json:"status"`
}

type JSONStatusOKWithDataResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type JSONErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
