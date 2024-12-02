package dto

type Response struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type SuccessResponse struct {
	Success string `json:"success" example:"true"`
}

// var TimeoutResponse = ErrorResponse{
// 	Message: "Your request has timed out",
// 	Error:   "timeout",
// }
