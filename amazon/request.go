package amazon

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Request is a request that amazon would send us via an API Gateway
type Request struct {
	events.APIGatewayProxyRequest
}

// NewRequest transforms an Amazon Gateway request into an IRequest, compatible with our
// controllers
func NewRequest(arg events.APIGatewayProxyRequest) (req *Request) {
	return &Request{arg}
}

// GetHeader returns a header in string format. If there is more than 1 of this type, a
// comma-separated version is returned
func (req *Request) GetHeader(key string) string {
	if header, ok := req.Headers[key]; ok {
		return header
	}

	if headers, ok := req.MultiValueHeaders[key]; ok {
		return strings.Join(headers, ",")
	}

	return ""
}

// GetBody just gets the raw body of the request
func (req *Request) GetBody() string {
	return req.Body
}

// GetBodyAs unmarshals the body to a given type. See `json.Unmarshal` for usage.
func (req *Request) GetBodyAs(arg interface{}) error {
	data := []byte(req.Body)
	return json.Unmarshal(data, arg)
}

// GetMethod returns the HTTP verb from the amazon request, guaranteed uppercase
func (req *Request) GetMethod() string {
	return strings.ToUpper(req.RequestContext.HTTPMethod)
}

// GetURLParam returns a url /path/:param to the user
func (req *Request) GetURLParam(arg string) string {
	if param, ok := req.PathParameters[arg]; ok {
		return param
	}

	return ""
}

// GetQueryParam returns a ?query=param to the user. If there is more than 1 of this type,
// a comma-separated version is returned
func (req *Request) GetQueryParam(key string) string {
	if query, ok := req.QueryStringParameters[key]; ok {
		return query
	}

	if query, ok := req.MultiValueQueryStringParameters[key]; ok {
		return strings.Join(query, ",")
	}

	return ""
}
