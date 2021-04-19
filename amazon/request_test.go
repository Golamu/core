package amazon_test

import (
	"fmt"
	"testing"

	"github.com/Golamu/core/amazon"
	"github.com/Golamu/core/http"
)

func TestRequestInterface(t *testing.T) {
	runThis := func(arg http.IRequest) {
		fmt.Printf("If this runs, it is compatible as an IRequest")
	}

	req := &amazon.Request{}
	runThis(req)
}
