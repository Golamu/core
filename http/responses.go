package http

// BaseResponse is a basic data structure that you can extend to keep a consistency across
// all your responses and give the user as much detail as you need
// NOTE: If you provide data, make sure the JSON tag is "data"
type BaseResponse struct {
	StatusCode int         `json:"statusCode"`
	RequestID  string      `json:"requestId,omitempty"`
	TimeTaken  int64       `json:"timeTaken,omitempty"`
	Messages   []string    `json:"messages"`
	Errors     []string    `json:"errors"`
	Data       interface{} `json:"data"`
}

// MessageResponse is the standard "It worked, here's a message" response to the user
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is the standard "It broke, here's a message" response to the user.
// ErrorResponse.Message is a user-readable string, and `Error` is used by the developer
// to determine what went wrong
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// ValidationErrorResponse is the type that is returned to the user when they've sent us a bad
// request that has failed validation
type ValidationErrorResponse struct {
	BaseResponse
	Validation map[string][]string `json:"data"`
}
