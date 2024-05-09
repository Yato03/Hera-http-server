package http

import (
	"fmt"
)

func OK(body string) Response {

	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     200,
		StatusText: "OK",
		Body:       body,
	}

	if body != "" {
		contentLength := fmt.Sprintf("%d", len(body))
		response.Headers = map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": contentLength,
		}
	}

	return response
}

func NOT_FOUND() Response {
	return Response{
		Protocol:   "HTTP/1.1",
		Status:     404,
		StatusText: "Not Found",
	}
}
