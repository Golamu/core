package http

import (
	"time"

	ohttp "github.com/Golamu/core/http"
)

// IRequest is a re-export from the Golamu/core/http package
type IRequest = ohttp.IRequest

// IResponse is a re-export from the Golamu/core/http package
type IResponse = ohttp.IResponse

// IContext is a re-export from the Golamu/core/http package
type IContext = ohttp.IContext

// Context is a basic, empty, no-effect HttpContext for use in unit testing controllers
type Context struct {
	req         IRequest
	res         IResponse
	doneChannel chan bool
	done        bool
	startTime   time.Time
}

// NewContext creates an http context object where you can define the request and response objects
func NewContext(req IRequest, res IResponse) *Context {
	var done = make(chan (bool), 1)
	return &Context{req, res, done, false, time.Now()}
}

// AddError wraps the response's AddError
func (ctx *Context) AddError(err ...string) {
	ctx.res.AddError(err...)
}

// GetDoneChannel returns the created done channel for testing
func (ctx *Context) GetDoneChannel() chan bool {
	return ctx.doneChannel
}

// IsDone returns whether or not something has signalled a completion
func (ctx *Context) IsDone() bool {
	return ctx.done
}

// Started returns the context's creation time
func (ctx *Context) Started() time.Time {
	return ctx.startTime
}

// Send sends the "done" message so we can reply to the client. Ignores multiple calls
func (ctx *Context) Send() {
	if ctx.done {
		return
	}

	ctx.doneChannel <- true
}

// AddMessage wraps the response's AddMessage
func (ctx *Context) AddMessage(msg ...string) {
	ctx.res.AddMessage(msg...)
}

// GetRequest gets the defined IRequest object initialized with NewTestContext
func (ctx *Context) GetRequest() IRequest {
	return ctx.req
}

// GetResponse gets the defined IResponse object initialized with NewTestContext
func (ctx *Context) GetResponse() IResponse {
	return ctx.res
}
