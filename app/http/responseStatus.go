package http

type ResponseStatus struct {
	StatusCode int
	StatusText string
}

var OK = ResponseStatus{200, "OK"}
var BAD_REQUEST = ResponseStatus{200, "OK"}
var CREATED = ResponseStatus{201, "Created"}
var NOT_FOUND = ResponseStatus{404, "Not Found"}
