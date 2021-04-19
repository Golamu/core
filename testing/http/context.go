package http

import (
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
	req IRequest
	res IResponse
}

// NewContext creates an http context object where you can define the request and response objects
func NewContext(req IRequest, res IResponse) *Context {
	return &Context{req: req, res: res}
}

// GetRequest gets the defined IRequest object initialized with NewTestContext
func (ctx *Context) GetRequest() IRequest {
	return ctx.req
}

// GetResponse gets the defined IResponse object initialized with NewTestContext
func (ctx *Context) GetResponse() IResponse {
	return ctx.res
}
