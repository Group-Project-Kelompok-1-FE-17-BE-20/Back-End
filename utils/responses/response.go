package responses

type MapResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type DataResponse struct {
	Code       int         `json:"code,omitempty"`
	Status     string      `json:"status,omitempty"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func WebResponse(code int, message string, data interface{}) MapResponse {
	return MapResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func ResponseFormat(code int, status, message string, data interface{}, pagination interface{}) DataResponse {
	result := DataResponse{
		Code:       code,
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	return result
}
