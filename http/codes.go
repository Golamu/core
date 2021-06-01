package http

// NOTE: This file was scraped from the MDN so you don't have to go to it yourself

const (
	// Continue - This interim response indicates that everything so far is OK and that the
	// client should continue the request or ignore the response if the request is already
	// finished.
	Continue = 100

	// SwitchingProtocol - This code is sent in response to an Upgrade request header from
	// the client and indicates the protocol the server is switching to.
	SwitchingProtocol = 101

	// Processing - This code indicates that the server has received and is processing the
	// request but no response is available yet.
	Processing = 102

	// EarlyHints - This status code is primarily intended to be used with the Link header
	// letting the user agent start preloading resources while the server prepares a response.
	EarlyHints = 103

	// OK - The request has succeeded. The meaning of the success depends on the HTTP method:
	// GET: The resource has been fetched and is transmitted in the msg body.
	// HEAD: The entity headers are in the msg body.
	// PUT/POST: The resource describing the result of the action is transmitted in the msg body.
	// TRACE: The msg body contains the request message as received by the server.
	OK = 200

	// Created - The request has succeeded and a new resource has been created as a result.
	// This is typically the response sent after POST requests or some PUT requests.
	Created = 201

	// Accepted - The request has been received but not yet acted upon. It is noncommittal
	// since there is no way in HTTP to later send an asynchronous response indicating the
	// outcome of the request. It is intended for cases where another process or server
	// handles the request or for batch processing.
	Accepted = 202

	// NonAuthoritativeInformation - This response code means the returned meta-information
	// is not exactly the same as is available from the origin server but is collected from
	// a local or a third-party copy. This is mostly used for mirrors or backups of another
	// resource. Except for that specific case the "200 OK" response is preferred to this
	// status.
	NonAuthoritativeInformation = 203

	// NoContent - There is no content to send for this request but the headers may be
	// useful. The user-agent may update its cached headers for this resource with the new
	// ones.
	NoContent = 204

	// ResetContent - Tells the user-agent to reset the document which sent this request.
	ResetContent = 205

	// PartialContent - This response code is used when the Range header is sent from the
	// client to request only part of a resource.
	PartialContent = 206

	// MultiStatus - Conveys information about multiple resources for situations where
	// multiple status codes might be appropriate.
	MultiStatus = 207

	// AlreadyReported - Used inside a <dav:propstat> response element to avoid repeatedly
	// enumerating the internal members of multiple bindings to the same collection.
	AlreadyReported = 208

	// IMUsed - The server has fulfilled a GET request for the resource and the response is
	// a representation of the result of one or more instance-manipulations applied to the
	// current instance.
	IMUsed = 226

	// MultipleChoice - The request has more than one possible response. The user-agent or
	// user should choose one of them. (There is no standardized way of choosing one of the
	// responses but HTML links to the possibilities are recommended so the user can pick.)
	MultipleChoice = 300

	// MovedPermanently - The URL of the requested resource has been changed permanently.
	// The new URL is given in the response.
	MovedPermanently = 301

	// Found - This response code means that the URI of requested resource has been changed
	// temporarily. Further changes in the URI might be made in the future. Therefore this
	// same URI should be used by the client in future requests.
	Found = 302

	// SeeOther - The server sent this response to direct the client to get the requested
	// resource at another URI with a GET request.
	SeeOther = 303

	// NotModified - This is used for caching purposes. It tells the client that the
	// response has not been modified so the client can continue to use the same cached
	// version of the response.
	NotModified = 304

	// TemporaryRedirect - Defined in a previous version of the HTTP specification to
	// indicate that a requested response must be accessed by a proxy. It has been
	// deprecated due to security concerns regarding in-band configuration of a
	// proxy.
	TemporaryRedirect = 307

	// PermanentRedirect - This response code is no longer used; it is just reserved. It
	// was used in a previous version of the HTTP/1.1 specification.
	PermanentRedirect = 308

	// BadRequest - The server sends this response to direct the client to get the
	// requested resource at another URI with same method that was used in the prior request.
	// This has the same semantics as the 302 Found HTTP response code with the exception
	// that the user agent must not change the HTTP method used: If a POST was used in the
	// first request a POST must be used in the second request.
	BadRequest = 400

	// Unauthorized - This means that the resource is now permanently located at another
	// URI specified by the Location: HTTP Response header. This has the same semantics as
	// the 301 Moved Permanently HTTP response code with the exception that the user agent
	// must not change the HTTP method used: If a POST was used in the first request a POST
	// must be used in the second request.
	Unauthorized = 401

	// PaymentRequired - The server could not understand the request due to invalid syntax.
	PaymentRequired = 402

	// Forbidden - Although the HTTP standard specifies "unauthorized" semantically this
	// response means "unauthenticated". That is the client must authenticate itself to get
	// the requested response.
	Forbidden = 403

	// NotFound - This response code is reserved for future use. The initial aim for
	// creating this code was using it for digital payment systems however this status code
	// is used very rarely and no standard convention exists.
	NotFound = 404

	// MethodNotAllowed - The client does not have access rights to the content; that is it
	// is unauthorized so the server is refusing to give the requested resource. Unlike 401
	// the client's identity is known to the server.
	MethodNotAllowed = 405

	// NotAcceptable - The server can not find the requested resource. In the browser this
	// means the URL is not recognized. In an API this can also mean that the endpoint is
	// valid but the resource itself does not exist. Servers may also send this response
	// instead of 403 to hide the existence of a resource from an unauthorized client. This
	// response code is probably the most famous one due to its frequent occurrence on the web.
	NotAcceptable = 406

	// ProxyAuthenticationRequired - The request method is known by the server but has been
	// disabled and cannot be used. For example an API may forbid DELETE-ing a resource.
	// The two mandatory methods GET and HEAD must never be disabled and should not return
	// this error code.
	ProxyAuthenticationRequired = 407

	// RequestTimeout - This response is sent when the web server after performing server
	// driven content negotiation doesn't find any content that conforms to the criteria
	// given by the user agent.
	RequestTimeout = 408

	// Conflict - This is similar to 401 but authentication is needed to be done by a proxy.
	Conflict = 409

	// Gone - This response is sent on an idle connection by some servers even without any
	// previous request by the client. It means that the server would like to shut down
	// this unused connection. This response is used much more since some browsers like
	// Chrome Firefox 27+ or IE9 use HTTP pre-connection mechanisms to speed up surfing.
	// Also note that some servers merely shut down the connection without sending this message.
	Gone = 410

	// LengthRequired - This response is sent when a request conflicts with the current
	// state of the server.
	LengthRequired = 411

	// PreconditionFailed - This response is sent when the requested content has been
	// permanently deleted from server with no forwarding address. Clients are expected to
	// remove their caches and links to the resource. The HTTP specification intends this
	// status code to be used for "limited-time promotional services". APIs should not feel
	// compelled to indicate resources that have been deleted with this status code.
	PreconditionFailed = 412

	// PayloadTooLarge - Server rejected the request because the Content-Length header
	// field is not defined and the server requires it.
	PayloadTooLarge = 413

	// URITooLong - The client has indicated preconditions in its headers which the server
	// does not meet.
	URITooLong = 414

	// UnsupportedMediaType - Request entity is larger than limits defined by server; the
	// server might close the connection or return an Retry-After header field.
	UnsupportedMediaType = 415

	// RangeNotSatisfiable - The URI requested by the client is longer than the server is
	// willing to interpret.
	RangeNotSatisfiable = 416

	// ExpectationFailed - The media format of the requested data is not supported by the
	// server so the server is rejecting the request.
	ExpectationFailed = 417

	// Imateapot - The range specified by the Range header field in the request can't be
	// fulfilled; it's possible that the range is outside the size of the target URI's data.
	Imateapot = 418

	// MisdirectedRequest - This response code means the expectation indicated by the
	// Expect request header field can't be met by the server.
	MisdirectedRequest = 421

	// UnprocessableEntity - The server refuses the attempt to brew coffee with a teapot.
	UnprocessableEntity = 422

	// Locked - The request was directed at a server that is not able to produce a respons
	// . This can be sent by a server that is not configured to produce responses for the
	// combination of scheme and authority that are included in the request URI.
	Locked = 423

	// FailedDependency - The request was well-formed but was unable to be followed due to
	// semantic errors.
	FailedDependency = 424

	// TooEarly - The resource that is being accessed is locked.
	TooEarly = 425

	// UpgradeRequired - The request failed due to failure of a previous request.
	UpgradeRequired = 426

	// PreconditionRequired - Indicates that the server is unwilling to risk processing a
	// request that might be replayed.
	PreconditionRequired = 428

	// TooManyRequests - The server refuses to perform the request using the current
	// protocol but might be willing to do so after the client upgrades to a different
	// protocol. The server sends an Upgrade header in a 426 response to indicate the
	// required protocol(s).
	TooManyRequests = 429

	// RequestHeaderFieldsTooLarge - The origin server requires the request to be
	// conditional. This response is intended to prevent the 'lost update' problem where a
	// client GETs a resource's state modifies it and PUTs it back to the server when
	// meanwhile a third party has modified the state on the server leading to a conflict.
	RequestHeaderFieldsTooLarge = 431

	// UnavailableForLegalReasons - The user has sent too many requests in a given amount
	// of time ("rate limiting").
	UnavailableForLegalReasons = 451

	// InternalServerError - The server is unwilling to process the request because its
	// header fields are too large. The request may be resubmitted after reducing the size
	// of the request header fields.
	InternalServerError = 500

	// NotImplemented - The user-agent requested a resource that cannot legally be provided
	// such as a web page censored by a government.
	NotImplemented = 501

	// BadGateway - The server has encountered a situation it doesn't know how to handle.
	BadGateway = 502

	// ServiceUnavailable - The request method is not supported by the server and cannot be
	// handled. The only methods that servers are required to support (and therefore that
	// must not return this code) are GET and HEAD.
	ServiceUnavailable = 503

	// GatewayTimeout - This error response means that the server while working as a
	// gateway to get a response needed to handle the request got an invalid response.
	GatewayTimeout = 504

	// HTTPVersionNotSupported - The server is not ready to handle the request. Common
	// causes are a server that is down for maintenance or that is overloaded. Note that
	// together with this response a user-friendly page explaining the problem should be
	// sent. This responses should be used for temporary conditions and the Retry-After:
	// HTTP header should if possible contain the estimated time before the recovery of the
	// service. The webmaster must also take care about the caching-related headers that
	// are sent along with this response as these temporary condition responses should
	// usually not be cached.
	HTTPVersionNotSupported = 505

	// VariantAlsoNegotiates - This error response is given when the server is acting as a
	// gateway and cannot get a response in time.
	VariantAlsoNegotiates = 506

	// InsufficientStorage - The HTTP version used in the request is not supported by the server.
	InsufficientStorage = 507

	// LoopDetected - The server has an internal configuration error: the chosen variant
	// resource is configured to engage in transparent content negotiation itself and is
	// therefore not a proper end point in the negotiation process.
	LoopDetected = 508

	// NotExtended - The method could not be performed on the resource because the server
	// is unable to store the representation needed to successfully complete the request.
	NotExtended = 510

	// NetworkAuthenticationRequired - The server detected an infinite loop while
	// processing the request.
	NetworkAuthenticationRequired = 511
)
