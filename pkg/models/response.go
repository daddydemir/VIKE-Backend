package models

type Response struct {
	Success bool
	Message string
	Data    interface{}
}

func (r Response) SuccessResponse(data interface{}) Response {
	return Response{Success: true, Message: "İşlem başarılı.", Data: data}
}

func (r Response) ErrorResponse(message string) Response {
	return Response{Message: message}
}
