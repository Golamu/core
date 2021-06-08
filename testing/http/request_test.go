package http_test

import (
	"testing"

	base "github.com/Golamu/core/http"
	http "github.com/Golamu/core/testing/http"
)

func Test_BuildsRequest(t *testing.T) {
	confirm := func(ctx base.IRequest) {}
	data := &http.Request{}
	confirm(data)
}
