package http

import (
	"testing"

	"github.com/go-test/deep"
	ohttp "github.com/Golamu/core/http"
)

// Response is a request type that allows you to easily test an endpoint's response
// without having to build something yourself
type Response struct {
	headers map[string]string
	query   map[string]string
	code    int
	body    interface{}
	test    *testing.T
}

// NewResponse allows you to have a default response for your testing purposes
func NewResponse(test *testing.T) *Response {
	resp := Response{
		headers: make(map[string]string),
		query:   make(map[string]string),
		code:    200,
		test:    test,
	}

	return &resp
}

// SetHeader allows the endpoint to set a header for future testing
func (res *Response) SetHeader(key, value string) {
	res.headers[key] = value
}

// HeaderMatches yes
func (res *Response) HeaderMatches(key string, value string) *Response {
	header := res.GetHeader(key)
	if header == value {
		return res
	}

	res.test.Errorf("Expected header '%s' to be '%s', actually '%s'", key, value, header)
	return res
}

// GetHeader allows you to retrieve a specific header from the response
func (res *Response) GetHeader(key string) string {
	if header, ok := res.headers[key]; ok {
		return header
	}

	return ""
}

// SetBody allows the user to assign any value to the body
func (res *Response) SetBody(arg interface{}) error {
	res.body = arg
	return nil
}

// GetBody allows you to retrieve the body as it was assigned
func (res *Response) GetBody() interface{} {
	return res.body
}

// SetCode allows the user to assign an HTTP status code to the request
func (res *Response) SetCode(code int) {
	res.code = code
	return
}

// HasStatus asserts that the response has the given status code
func (res *Response) HasStatus(code int) *Response {
	if code == res.code {
		return res
	}

	res.test.Errorf("Expected status '%d', actually '%d'", code, res.code)

	return res
}

// HasMessage assert that the response has been assigned a specific message
func (res *Response) HasMessage(arg string) *Response {
	msg, ok := res.body.(ohttp.MessageResponse)
	if !ok {
		res.test.Errorf("Body is not of type MessageResponse")
		return res
	}

	if msg.Message != arg {
		res.test.Errorf("Messages do not match.\n\tExpected: %s\n---\n\n\tGot: %s", arg, msg.Message)
	}

	return res
}

// HasError asserts that the user has set a proper ErrorResponse
func (res *Response) HasError(err string, code int, msg string) *Response {
	body, ok := res.body.(ohttp.ErrorResponse)
	if !ok {
		res.test.Errorf("Body is not of type ErrorResponse")
		return res
	}

	errFormat := "%s do not match.\n\tExpected: %s\n----\n\n\tGot: %s"

	if body.Message != msg {
		res.test.Errorf(errFormat, "Messages", msg, body.Message)
	}

	if body.Error != err {
		res.test.Errorf(errFormat, "Errors", err, body.Error)
	}

	return res.HasStatus(code)
}

// BodyMatches assert that the body of the response matches a given value
func (res *Response) BodyMatches(arg interface{}) *Response {
	if diff := deep.Equal(res.body, arg); diff != nil {
		res.test.Error(diff)
		res.test.FailNow()
	}

	return res
}

// LogBody prints the formatted string version of the body
func (res *Response) LogBody() *Response {
	str, err := getPrettyBodyString(res.body)
	if err != nil {
		res.test.Logf("Unable to print body because: %s", err.Error())
		return res
	}

	res.test.Logf("Body:\n----\n%s\n----\n", str)

	return res
}

// SetMessage allows the user to respond with a message
func (res *Response) SetMessage(msg string) error {
	resp := ohttp.MessageResponse{Message: msg}
	return res.SetBody(resp)
}

// SetError allows the user to quickly and easily respond with an error
func (res *Response) SetError(code int, message, errorMessage string) error {
	res.code = code
	resp := ohttp.ErrorResponse{Error: errorMessage, Message: message}
	return res.SetBody(resp)
}

// ForContext provides an easy type-cast for IContext
func (res *Response) ForContext() IResponse {
	return res
}
