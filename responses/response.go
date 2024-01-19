package responses

type MapResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(code int, message string, data interface{}) MapResponse {
	return MapResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
