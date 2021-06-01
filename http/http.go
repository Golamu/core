package http

import "time"

// IRequest represents the basics an incoming request type needs to provide to process an incoming
// request, independent of any endpoint provider
type IRequest interface {
	GetHeader(key string) string
	GetBody() string
	GetBodyAs(arg interface{}) error
	GetMethod() string
	GetID() string
	GetURLParam(arg string) string
	GetQueryParam(param string) string
}

// IResponse represents the basics needed to process a response to the user independent of any
// endpoint provider
type IResponse interface {
	SetHeader(key, value string) error
	GetHeader(key string) string
	SetBody(arg interface{}) error
	SetCode(code int) error

	AddError(err ...string) error
	AddMessage(msg ...string) error
	SetData(arg interface{}) error

	// Finish should set a flag that freezes the request in place, refusing other updates
	Finish() error
}

// IContext represents the type that all controllers must support in order to process an HTTP
// request
type IContext interface {
	GetRequest() IRequest
	GetResponse() IResponse

	AddError(err ...string)
	AddMessage(message ...string)

	SetData(arg map[string]interface{})
	SetDataAt(key string, arg interface{})
	GetDataAt(key string) interface{}
	GetData() map[string]interface{}

	// Return this request to the client
	Send()
	IsDone() bool
	GetDoneChannel() chan bool

	Started() time.Time
}
