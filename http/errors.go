package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/Golamu/core"
)

// InternalServerException takes an IContext pointer, assigns the "something bad happened" error,
// and then sets the error code for you.
func InternalServerException(ctx IContext) {
	msg := "Uh oh! Something went wrong with our server. Contact support if it persists."
	error := "An internal error occurred"
	ctx.GetResponse().SetError(500, msg, error)
}

// FailedValidationException receives a set of validation errors and transforms it into something
// that is a little easier to parse for clients
func FailedValidationException(ctx IContext, errors validator.ValidationErrors) {
	msg := FailedValidationExceptionMessage
	err := FailedValidationExceptionError
	valid := core.ValidationErrorDict(errors)

	respBody := ValidationErrorResponse{
		Error:      err,
		Message:    msg,
		Validation: valid,
	}

	resp := ctx.GetResponse()
	resp.SetBody(respBody)
	resp.SetCode(400)
}

// ConflictException automatically writes the code and passes on the error
// and message to the response body
func ConflictException(ctx IContext, msg string, err string) {
	ctx.GetResponse().SetError(409, msg, err)
}
