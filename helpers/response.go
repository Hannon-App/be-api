package helpers

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

type FindAllMapResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	NextPage bool        `json:"next"`
	Data     interface{} `json:"data,omitempty"`
}

func FindAllWebResponse(code int, message string, data interface{}, nextPage bool) FindAllMapResponse {
	return FindAllMapResponse{
		Code:     code,
		Message:  message,
		Data:     data,
		NextPage: nextPage,
	}
}
