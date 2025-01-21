package helper

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func FailedResponse(message string) Response {
	return Response{
		Message: message,
	}
}

func AuthResponse(message, token string) Response {
	return Response{
		Success: true,
		Message: message,
		Token:   token,
	}
}
