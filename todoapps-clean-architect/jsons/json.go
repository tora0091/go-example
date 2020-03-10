package jsons

type JSONStatusOKResponse struct {
	Status int `json:"Status"`
}

type JSONStatusOKWithDataResponse struct {
	Status int         `json:"Status"`
	Data   interface{} `json:"Data"`
}

type JSONErrorResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}
