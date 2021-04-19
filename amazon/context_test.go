package amazon_test

import (
	"fmt"
	"testing"

	"github.com/Golamu/core/amazon"
	"github.com/Golamu/core/http"
)

func TestContextInterface(t *testing.T) {

	runThis := func(arg http.IContext) {
		fmt.Printf("If this runs, it is compatible as an IContext")
	}

	ctx := &amazon.Context{}
	runThis(ctx)
}
