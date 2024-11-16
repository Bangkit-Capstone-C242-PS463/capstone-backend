package dto

type Response struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

var SuccessResponse = Response{
	Message: "success",
}

// var TimeoutResponse = ErrorResponse{
// 	Message: "Your request has timed out",
// 	Error:   "timeout",
// }
