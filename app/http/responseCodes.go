package http

import (
	"fmt"
)

func OK(body string, request Request) Response {

	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     200,
		StatusText: "OK",
		Body:       body,
		Headers:    map[string]string{},
	}

	//Content-Length
	if body != "" {
		contentLength := fmt.Sprintf("%d", len(body))
		response.Headers["Content-Type"] = "text/plain"
		response.Headers["Content-Length"] = contentLength
	}

	//Accept-Encoding
	if request.Headers["Accept-Encoding"] != "" && request.Headers["Accept-Encoding"] == "gzip" {
		response.Headers["Content-Encoding"] = "gzip"
		//response.Body = Gzip(response.Body)
	}

	return response
}

func OK_FILE(body string) Response {

	contentLength := fmt.Sprintf("%d", len(body))

	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     200,
		StatusText: "OK",
		Body:       body,
		Headers: map[string]string{
			"Content-Type":   "application/octet-stream",
			"Content-Length": contentLength,
		},
	}

	return response
}

func CREATED() Response {
	return Response{
		Protocol:   "HTTP/1.1",
		Status:     201,
		StatusText: "Created",
	}
}

func BAD_REQUEST(body string) Response {
	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     400,
		StatusText: "Bad Request",
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
