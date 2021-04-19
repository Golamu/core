package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/Golamu/core/http"
)

// Validate takes the context and a DTO _pointer_, and validates the body. Automatically
// pre-processes the response if the request fails validation.
func Validate(ctx http.IContext, vld *validator.Validate, arg interface{}) bool {

	if err := ctx.GetRequest().GetBodyAs(arg); err != nil {
		http.InternalServerException(ctx)
		return false
	}

	switch err := vld.Struct(arg); err.(type) {
	case validator.ValidationErrors:
		http.FailedValidationException(ctx, err.(validator.ValidationErrors))
		return false

	case nil:
		break

	default:
		http.InternalServerException(ctx)
		return false
	}

	return true
}
