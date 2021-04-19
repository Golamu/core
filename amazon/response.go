package amazon

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/Golamu/core/http"
)

// Response is the type that is compatible with both our controllers, and Amazon
type Response struct {
	events.APIGatewayProxyResponse
}

// NewResponse creates a new response with sensible defaults
func NewResponse() *Response {
	evt := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"content-type": "application/json",
		},
	}
	return &Response{evt}
}

// SetHeader sets the `key` HTTP header on the response object to `value`
func (res *Response) SetHeader(key, value string) {
	res.Headers[key] = value
}

// GetHeader returns the `key` HTTP header that's already set on the response object
func (res *Response) GetHeader(key string) string {
	return res.Headers[key]
}

// SetBody overwrites the body with the json marshaled version of arg
func (res *Response) SetBody(arg interface{}) (err error) {
	var body []byte

	if body, err = json.Marshal(arg); err != nil {
		return
	}

	res.Body = string(body)

	return nil
}

// SetCode sets the StatusCode on the amazon response
func (res *Response) SetCode(num int) {
	res.StatusCode = num
}

// SetMessage overwrites the body with a MessageResponse type
func (res *Response) SetMessage(msg string) error {
	resp := http.MessageResponse{Message: msg}
	return res.SetBody(resp)
}

// SetError overwrites the body with an ErrorResponse type
func (res *Response) SetError(code int, message, errorMessage string) error {
	resp := http.ErrorResponse{
		Message: message,
		Error:   errorMessage,
	}

	err := res.SetBody(resp)
	if err != nil {
		return err
	}

	res.SetCode(code)
	return nil
}

// Respond gives the APIGatewayProxyResponse for use in the endpoint handlers
func (res *Response) Respond() events.APIGatewayProxyResponse {
	return res.APIGatewayProxyResponse
}
