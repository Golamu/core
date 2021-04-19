package core

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationErrorDict generats a map[string]string based on the errors in validator.ValidationErrors
func ValidationErrorDict(errors validator.ValidationErrors) (valid map[string][]string) {
	valid = make(map[string][]string)

	for _, err := range errors {
		if valid[err.Field()] == nil {
			valid[err.Field()] = make([]string, 0)
		}

		msg := fmt.Sprintf(
			"%s failed with value \"%s\", (%s)",
			err.ActualTag(),
			err.Param(),
			err.Type(),
		)

		valid[err.Field()] = append(valid[err.Field()], msg)
	}

	return
}
