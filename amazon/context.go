package amazon

import (
	"time"

	"github.com/Golamu/core/http"
	"github.com/aws/aws-lambda-go/events"
)

// Context is the Amazon context created by a handler to pass into a controller
// handler
type Context struct {
	http.IContext
	request     *Request
	response    *Response
	started     time.Time
	doneChannel chan bool
	done        bool
}

// NewContext creates the request, response, and context objects for use in
// a controller for this API
func NewContext(req events.APIGatewayProxyRequest) *Context {
	request := NewRequest(req)
	response := NewResponse(request)

	ctx := &Context{
		request:     request,
		response:    response,
		started:     time.Now().UTC(),
		doneChannel: make(chan bool, 1),
		done:        false,
	}

	return ctx
}

// GetRequest gets the context's request object
func (ctx *Context) GetRequest() http.IRequest {
	return ctx.request
}

// GetResponse gets the context's response object
func (ctx *Context) GetResponse() http.IResponse {
	return ctx.response
}

// Respond gives the APIGatewayProxyResponse for use in the endpoint handlers
func (ctx *Context) Respond() events.APIGatewayProxyResponse {
	return ctx.response.APIGatewayProxyResponse
}

// GetDoneChannel returns the "is it done" channel for an amazon request
func (ctx *Context) GetDoneChannel() chan bool {
	return ctx.doneChannel
}

// AddError is an easy way to add a messages to the body of the response
func (ctx *Context) AddError(msgs ...string) {
	ctx.response.AddError(msgs...)
}

// Send sends the "done" message so we can reply to the client. Ignores multiple calls
func (ctx *Context) Send() {
	if ctx.done {
		return
	}

	ctx.doneChannel <- true
}
