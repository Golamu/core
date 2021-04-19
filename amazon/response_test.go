package amazon_test

import (
	"fmt"
	"testing"

	"github.com/Golamu/core/amazon"
	"github.com/Golamu/core/http"
)

func TestResponseInterface(t *testing.T) {
	runThis := func(arg http.IResponse) {
		fmt.Printf("If this runs, it is compatible as an IResponse")
	}

	res := &amazon.Response{}
	runThis(res)
}
