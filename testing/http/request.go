package http

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
)

// Request is a request type that allows you to build a body for
// endpoint testing without having to hook up an HTTP library
type Request struct {
	headers map[string]string
	query   map[string]string
	params  map[string]string
	body    string
	verb    string
	id      string
	test    *testing.T
}

// NewRequest prepares a start IRequest so you can integration test endpoints
func NewRequest(t *testing.T) *Request {
	id, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	req := Request{
		id:      id.String(),
		headers: make(map[string]string),
		query:   make(map[string]string),
		test:    t,
	}

	req.SetHeader("content-type", "application/json")

	return &req
}

// GetHeader returns the header for key
func (req *Request) GetHeader(key string) string {
	header := strings.ToLower(key)
	return req.headers[header]
}

// SetHeader assigns the header key with the value val. If you want multiple header values,
// pass a comma separated string
func (req *Request) SetHeader(key string, val string) *Request {
	header := strings.ToLower(key)
	req.headers[header] = val
	return req
}

// SetMultiHeader allows you to pass an array of strings to a single header
func (req *Request) SetMultiHeader(key string, val []string) *Request {
	return req.SetHeader(key, strings.Join(val, ","))
}

// GetBody returns the string body of this request
func (req *Request) GetBody() string {
	return req.body
}

// SetBody allows you to assign any kind of "valid" request body for HTTP
// []byte is assumed to be a string-like value, and anything that's not a
// primitive is marshalled as JSON
func (req *Request) SetBody(arg interface{}) (*Request, error) {
	body, err := getBodyString(arg)
	req.body = body
	if err != nil {
		req.test.Errorf("Unable to set body: %s", err.Error())
	}
	return req, err
}

// GetBodyAs is implemented to allow the consumer to transform the request from JSON to a type
// of their choosing
func (req *Request) GetBodyAs(arg interface{}) error {
	data := []byte(req.body)
	return json.Unmarshal(data, arg)
}

// GetMethod returns the all-caps HTTP verb
func (req *Request) GetMethod() string {
	return req.verb
}

// SetMethod capitalizes arg and sets the HTTP verb for this request
func (req *Request) SetMethod(arg string) *Request {
	req.verb = strings.ToUpper(arg)
	return req
}

// GetURLParam returns a case-sensitive parameter of the URL
func (req *Request) GetURLParam(arg string) string {
	return req.params[arg]
}

// SetURLParam sets a case-sensitive URL param to arg
func (req *Request) SetURLParam(key string, arg string) *Request {
	req.params[key] = arg
	return req
}

// GetQueryParam returns a case-sensitive query parameter to a string
func (req *Request) GetQueryParam(param string) string {
	return req.query[param]
}

// SetQueryParam allows you to set a single-value query parameter to arg
func (req *Request) SetQueryParam(key string, arg string) *Request {
	req.query[key] = arg
	return req
}

// SetMultiQueryParam allows you to set an array of strings to the request's query param
func (req *Request) SetMultiQueryParam(key string, arg []string) *Request {
	return req.SetQueryParam(key, strings.Join(arg, ","))
}

// GetID gets the generated UUID for this test request
func (req *Request) GetID() string {
	return req.id
}

// ForContext provides a type translation for creating test contexts
func (req *Request) ForContext() IRequest {
	return req
}
