package http_test

import (
	"testing"

	base "github.com/Golamu/core/http"
	http "github.com/Golamu/core/testing/http"
)

func Test_BuildsContext(t *testing.T) {
	confirm := func(ctx base.IContext) {}
	data := &http.Context{}
	confirm(data)
}
