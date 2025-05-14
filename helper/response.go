package helper

type ResponseData struct {
	Data any `json:"data"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string `json:"message"`
}