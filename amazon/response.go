package amazon

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/Golamu/core/http"
	"github.com/aws/aws-lambda-go/events"
)

// Response is the type that is compatible with both our controllers, and Amazon
type Response struct {
	events.APIGatewayProxyResponse
	requestID string
	messages  []string
	errors    []string
	data      interface{}
	done      bool
	started   time.Time
}

// NewResponse creates a new response with sensible defaults, including the BaseResponse from
// this package. Note if you reassign
func NewResponse(req *Request) *Response {
	verb := req.GetMethod()
	code := http.GetVerbStatus(strings.ToUpper(verb))

	evt := events.APIGatewayProxyResponse{
		StatusCode:      code,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"content-type": "application/json",
		},
	}

	resp := &Response{
		evt,
		req.GetID(),
		make([]string, 0),
		make([]string, 0),
		nil,
		false,
		time.Now().UTC(),
	}

	return resp
}

// AddError adds a new error to the response ONLY WHEN custom data has never been set
func (res *Response) AddError(errs ...string) error {
	if res.done {
		return errors.New("Cannot add errors to a finished response")
	}

	for _, err := range errs {
		res.errors = append(res.errors, err)
	}

	return nil
}

// AddMessage adds a new message to the response ONLY WHEN custom data has never been set
func (res *Response) AddMessage(errs ...string) error {
	if res.done {
		return errors.New("Cannot add messages to a finished response")
	}

	for _, err := range errs {
		res.messages = append(res.messages, err)
	}

	return nil
}

// SetHeader sets the `key` HTTP header on the response object to `value`
func (res *Response) SetHeader(key, value string) error {
	if res.done {
		return errors.New("Unable to set a header on a finished response")
	}

	res.Headers[key] = value

	return nil
}

// GetHeader returns the `key` HTTP header that's already set on the response object
func (res *Response) GetHeader(key string) string {
	return res.Headers[key]
}

// SetBody overwrites the body with your own type. It will disable AddMessage and AddError if you
// use this and pass anything except BaseResponse, so use SetData instead if you are not
// looking to completely replace this
func (res *Response) SetBody(arg interface{}) (err error) {

	body, err := json.Marshal(arg)
	if err != nil {
		return err
	}

	res.Body = string(body)

	return nil
}

// SetData sets the data to be marshalled at the end of the request
func (res *Response) SetData(arg interface{}) error {
	if res.done {
		return errors.New("Unable to set data on finished response")
	}

	res.data = arg

	return nil
}

// SetCode sets the StatusCode on the amazon response
func (res *Response) SetCode(num int) error {
	if res.done {
		return errors.New("Unable to set code on finished response")
	}

	res.StatusCode = num
	return nil
}

// Finish locks the response so that nobody can alter it after it has been finished
func (res *Response) Finish() (err error) {
	if res.done {
		return
	}

	res.done = true

	if res.Body != "" {
		return
	}

	took := time.Since(res.started).Milliseconds()

	resp := http.BaseResponse{
		StatusCode: res.StatusCode,
		RequestID:  res.requestID,
		TimeTaken:  took,
		Data:       res.data,
		Messages:   res.messages,
		Errors:     res.errors,
	}

	var body []byte
	if body, err = json.Marshal(resp); err != nil {
		return err
	}

	res.Body = string(body)

	return nil
}
