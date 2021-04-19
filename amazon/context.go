package amazon

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/Golamu/core/http"
)

// Context is the Amazon context created by a handler to pass into a controller
// handler
type Context struct {
	Request  *Request
	Response *Response
}

// NewContext creates the request, response, and context objects for use in
// a controller for this API
func NewContext(req events.APIGatewayProxyRequest) *Context {
	request := NewRequest(req)
	response := NewResponse()

	return &Context{
		Request:  request,
		Response: response,
	}
}

// GetRequest gets the context's request object
func (ctx *Context) GetRequest() http.IRequest {
	return ctx.Request
}

// GetResponse gets the context's response object
func (ctx *Context) GetResponse() http.IResponse {
	return ctx.Response
}
