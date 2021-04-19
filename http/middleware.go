package http

// MiddlewareContext provides all the data necessary for middleware to process
// and handle a request lifecycle
type MiddlewareContext struct {
	Request  IRequest
	Response IResponse
	Data     map[string]interface{}
}

// Interceptor is the type that helps you re-process a response to a user, adding extra
// fields, setting headers, etc.
type Interceptor func(ctx *MiddlewareContext) bool
