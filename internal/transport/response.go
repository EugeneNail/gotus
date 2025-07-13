package transport

type Response struct {
	Status  int
	Message string
	Data    any
	Writer  ResponseWriter
}

func NewJsonResponse(status int, message string, data any) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
		Writer:  JsonResponseWriter{},
	}
}

func NewPlainResponse(status int, data string) *Response {
	return &Response{
		Status: status,
		Data:   data,
		Writer: PlainResponseWriter{},
	}
}
