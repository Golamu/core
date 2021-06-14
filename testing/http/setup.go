package http

import "testing"

// SetupCall takes your testing instance, and generates a new Context, Request, Response set
// for you to test your routes with
func SetupCall(t *testing.T) (*Context, *Request, *Response) {
	req := NewRequest(t)
	res := NewResponse(t)
	ctx := NewContext(req, res)
	return ctx, req, res
}
