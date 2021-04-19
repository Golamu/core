package http

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
	Message    string              `json:"message"`
	Error      string              `json:"error"`
	Validation map[string][]string `json:"validation"`
}
