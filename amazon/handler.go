package amazon

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// ProxyRequest is the incoming request from Amazon
type ProxyRequest = events.APIGatewayProxyRequest

// ProxyResponse is the outgoing response from Amazon
type ProxyResponse = events.APIGatewayProxyResponse

// Endpoint is a function that can handle a request / response cycle
type Endpoint func(ProxyRequest) (ProxyResponse, error)

// HTTPEndpoint is an HTTP Endpoint that takes just a context, and manages the rest itself
type HTTPEndpoint func(ctx *Context) error

// Start a Lambda function, auto-wrapping
func Start(pt Endpoint) {
	lambda.Start(pt)
}

// StartHTTP allows you to start an HTTP endpoint with the Golamu HTTP context and pass it
// into a function. This automatically creates a context and responds for you.
func StartHTTP(handler HTTPEndpoint) {
	Start(func(req ProxyRequest) (ProxyResponse, error) {
		ctx := NewContext(req)
		err := handler(ctx)
		return ctx.Response.Respond(), err
	})
}
