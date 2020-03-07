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

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Job       string    `json:"job"`
	CreatedAt time.Time `json:"created_ad"`
	UpdatedAt time.Time `json:"updated_ad"`
}

type Users []User

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
