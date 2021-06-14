package http

import (
	"github.com/Golamu/core"
	"github.com/go-playground/validator/v10"
)

// InternalServerException takes an IContext pointer, assigns the "something bad happened" error,
// and then sets the error code for you.
func InternalServerException(ctx IContext) {
	msg := "Uh oh! Something went wrong with our server. Contact support if it persists."
	err := "An internal error occurred"

	res := ctx.GetResponse()
	res.AddError(err)
	res.AddMessage(msg)
	res.SetCode(InternalServerError)
}

// FailedValidationException receives a set of validation errors and transforms it into something
// that is a little easier to parse for clients
func FailedValidationException(ctx IContext, errors validator.ValidationErrors) {
	valid := core.ValidationErrorDict(errors)

	base := BaseResponse{
		Errors:     []string{FailedValidationExceptionError},
		Messages:   []string{FailedValidationExceptionMessage},
		StatusCode: 400,
	}

	respBody := ValidationErrorResponse{
		BaseResponse: base,
		Validation:   valid,
	}

	resp := ctx.GetResponse()
	resp.SetBody(respBody)
	resp.SetCode(400)
}

// ConflictException automatically writes the code and passes on the error
// and message to the response body
func ConflictException(ctx IContext, msg string, err string) {
	res := ctx.GetResponse()
	res.AddError(err)
	res.AddMessage(msg)
	res.SetCode(409)
}

// UnauthorizedException provides a default exception for unauthorized access attempts
func UnauthorizedException(ctx IContext) {
	res := ctx.GetResponse()
	res.SetCode(Unauthorized)
	res.AddError("You are not authorized to access this resource")
	res.AddError("Verify you have a valid auth token, or that you have permission to access this")
}

// UnauthorizedExceptionCauseBy lets you create an unauthorized exception with a specific message
func UnauthorizedExceptionCauseBy(ctx IContext, msg string) {
	res := ctx.GetResponse()
	res.AddError("You are not authorized to access this resource")
	res.AddMessage(msg)
	res.SetCode(Unauthorized)
}
