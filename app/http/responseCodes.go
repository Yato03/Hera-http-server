package http

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/fileUtils"
)

func OK(body string, request Request) Response {

	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     200,
		StatusText: "OK",
		Body:       body,
		Headers:    map[string]string{},
	}

	//Accept-Encoding
	if request.Headers["Accept-Encoding"] != "" {
		encondings := strings.Split(request.Headers["Accept-Encoding"], ",")

		for _, encoding := range encondings {
			if strings.TrimSpace(encoding) == "gzip" {
				response.Headers["Content-Encoding"] = "gzip"
				encodedText, err := fileUtils.Gzip(response.Body)
				if err != nil {
					return BAD_REQUEST("Error compressing response")
				}
				response.Body = encodedText
			}
		}
	}

	//Content-Length
	if response.Body != "" {
		contentLength := fmt.Sprintf("%d", len(response.Body))
		response.Headers["Content-Type"] = "text/plain"
		response.Headers["Content-Length"] = contentLength
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
