package http

import "strings"

// DefaultCodeMap is the mapping of HTTP verbs (all caps) to
// HTTP Status codes on a successful response. Feel free to remap
// this to your heart's desire
var DefaultCodeMap = map[string]int{
	"POST":    Created,
	"GET":     OK,
	"DELETE":  OK,
	"PUT":     OK,
	"PATCH":   OK,
	"OPTIONS": OK,
	"HEAD":    OK,
}

// GetVerbStatus allows you to pass a request VERB in for a status code
// POST -> CREATED
// GET -> OK
// PATCH -> OK
// DELETE -> OK
func GetVerbStatus(arg string) int {
	verb := strings.ToUpper(arg)
	code, ok := DefaultCodeMap[verb]
	if !ok {
		return 200
	}

	return code
}
