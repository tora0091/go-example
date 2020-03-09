package jsons

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
