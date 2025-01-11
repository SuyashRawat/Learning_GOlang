package utils

type Response struct{
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`

}
func ErrorResponse(str string) Response{
	return Response{
		Status: "Error",
		Message: str,

	}

}
func SuccessResponse(str string) Response{
	
	return Response{
		Status: "Sucesss",
		Message: str,
	}
}
func SuccessResponseWithData(message string, data interface{}) Response {
	return Response{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
}
