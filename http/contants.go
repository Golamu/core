package http

const (
	// FailedValidationExceptionMessage is the message that we likely show to users when
	// they have sent us something that failed to validate
	FailedValidationExceptionMessage = "Uh oh! We received some information we couldn't understand. " +
		"Please try again."

	// FailedValidationExceptionError is the message that we assign to error so that devs
	// can parse the error message easier
	FailedValidationExceptionError = "Request failed validation"
)
