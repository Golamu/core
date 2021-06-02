package http

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/go-test/deep"
)

// Response is a request type that allows you to easily test an endpoint's response
// without having to build something yourself
type Response struct {
	headers  map[string]string
	query    map[string]string
	code     int
	body     interface{}
	data     interface{}
	test     *testing.T
	errors   []string
	messages []string
	done     bool
}

// NewResponse allows you to have a default response for your testing purposes
func NewResponse(test *testing.T) *Response {
	resp := Response{
		headers: make(map[string]string),
		query:   make(map[string]string),
		code:    200,
		test:    test,
		errors:  make([]string, 1),
		done:    false,
	}

	return &resp
}

// AddError adds a string error message to the errors array
func (res *Response) AddError(errs ...string) error {
	for _, msg := range errs {
		res.errors = append(res.errors, msg)
	}

	return nil
}

// AddMessage adds a string message message to the messages array
func (res *Response) AddMessage(errs ...string) error {
	if res.done {
		return errors.New("Response has already been completed, cannot add messages")
	}

	for _, msg := range errs {
		res.messages = append(res.messages, msg)
	}

	return nil
}

// SetHeader allows the endpoint to set a header for future testing
func (res *Response) SetHeader(key, value string) error {
	if res.done {
		return errors.New("Response has already been completed, cannot add messages")
	}

	res.headers[key] = value
	return nil
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

// SetBody allows the user to assign any value to the body. We marshal the argument to verify
// that the endpoint is responding properly to marshalling errors
func (res *Response) SetBody(arg interface{}) (err error) {
	if res.done {
		return errors.New("Unable to set body because response has already been finished")
	}

	res.body = arg

	_, err = json.Marshal(arg)

	return
}

// SetData simply sets the data property for this object
func (res *Response) SetData(arg interface{}) (err error) {
	res.data = arg
	return
}

// GetBody allows you to retrieve the body as it was assigned
func (res *Response) GetBody() interface{} {
	return res.body
}

// SetCode allows the user to assign an HTTP status code to the request
func (res *Response) SetCode(code int) error {
	if res.done {
		return errors.New("Unable to set code because response has already been finished")
	}

	res.code = code
	return nil
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
	if arg == "" {
		return res
	}

	for _, msg := range res.messages {
		if msg == arg {
			return res
		}
	}

	res.test.Errorf("Messages not found.\n\tExpected: %s\n---\n\n", arg)
	return res
}

// HasError asserts that the user has set a proper ErrorResponse
func (res *Response) HasError(err string, code int, msg string) *Response {
	if err == "" {
		return res
	}

	for _, msg := range res.errors {
		if msg == err {
			return res.HasMessage(msg).HasStatus(code)
		}
	}

	res.test.Errorf("Errors not found.\n\tExpected: %s\n---\n\n", err)
	return res
}

// DataMatches asserts that the body of the response matches a given value
func (res *Response) DataMatches(arg interface{}) *Response {
	if diff := deep.Equal(res.data, arg); diff != nil {
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

// Finish sets the "done" flag
func (res *Response) Finish() error {
	res.done = true
	return nil
}

// ForContext provides an easy type-cast for IContext
func (res *Response) ForContext() IResponse {
	return res
}
