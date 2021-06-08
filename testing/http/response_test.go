package http_test

import (
	"testing"

	base "github.com/Golamu/core/http"
	http "github.com/Golamu/core/testing/http"
)

func Test_BuildsResponse(t *testing.T) {
	confirm := func(ctx base.IResponse) {}
	data := &http.Response{}
	confirm(data)
}
