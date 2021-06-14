package amazon

import (
	"context"
	"time"

	"github.com/Golamu/core/http"
	"github.com/aws/aws-lambda-go/events"
)

// Context is the Amazon context created by a handler to pass into a controller
// handler
type Context struct {
	request     *Request
	response    *Response
	started     time.Time
	doneChannel chan bool
	done        bool
	ctx         context.Context
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
		ctx:         context.Background(),
	}

	return ctx
}

// IsDone returns whether or not the response has been finalized
func (ctx *Context) IsDone() bool {
	return ctx.done
}

// Context returns this "request context's" call contex
func (ctx *Context) Context() context.Context {
	return ctx.ctx
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

// AddMessage is a binding to the response's "AddMessage" method
func (ctx *Context) AddMessage(arg ...string) {
	ctx.GetResponse().AddMessage(arg...)
}

// Send sends the "done" message so we can reply to the client. Ignores multiple calls
func (ctx *Context) Send() {
	if ctx.done {
		return
	}

	ctx.doneChannel <- true
}

// Started returns the time that this context was initialized
func (ctx *Context) Started() time.Time {
	return ctx.started
}
