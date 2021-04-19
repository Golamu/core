package http

// IRequest represents the basics an incoming request type needs to provide to process an incoming
// request, independent of any endpoint provider
type IRequest interface {
	GetHeader(key string) string
	GetBody() string
	GetBodyAs(arg interface{}) error
	GetMethod() string
	GetURLParam(arg string) string
	GetQueryParam(param string) string
}

// IResponse represents the basics needed to process a response to the user independent of any
// endpoint provider
type IResponse interface {
	SetHeader(key, value string)
	GetHeader(key string) string
	SetBody(arg interface{}) error
	SetCode(code int)
	SetMessage(msg string) error
	SetError(code int, message, errorMessage string) error
}

// IContext represents the type that all controllers must support in order to process an HTTP
// request
type IContext interface {
	GetRequest() IRequest
	GetResponse() IResponse
}
