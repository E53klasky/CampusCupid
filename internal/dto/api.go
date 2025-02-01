package dto

// ApiResponse is a standard response from the service.
type ApiResponse[T any] struct {
	Status bool      `json:"status"`
	Result *T        `json:"result"`
	Error  *ApiError `json:"error"`
}

// ApiError is a standard error response from the service.
type ApiError struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"http_status_code"`
}

// ErrorResponse is a standard error response from handler methods.
type ErrorResponse struct {
	Code           int
	Message        string
	HttpStatusCode int
}
