package helper

type EmptyObject []string

type ResponseSuccess struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

type ResponseFail struct {
	Status   string `json:"status" default:"fail"`
	Data     any    `json:"data"`
	DataFail any    `json:"data_fail"`
	Message  string `json:"message"`
}

// ApiResponseSuccess is a function to return response success
func ApiResponseSuccess(status string, data any) *ResponseSuccess {
	//return empty string if data is nil
	if isNilFixed(data) {
		data = EmptyObject{}
	}
	return &ResponseSuccess{
		Status: status,
		Data:   data,
	}
}

// ApiResponseError is a function to return response error
func ApiResponseError[T any](message T) *ResponseError {
	return &ResponseError{
		Status:  "error",
		Message: message,
	}
}

// ApiResponseFail is a function to return response fail
func ApiResponseFail(data any, message string) *ResponseFail {
	return &ResponseFail{
		Status:   "fail",
		Data:     data,
		DataFail: data,
		Message:  message,
	}
}
