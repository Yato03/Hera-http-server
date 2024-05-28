package http

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/fileUtils"
)

type Response struct {
	Protocol   string
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
}

func TextPlain(body string, request Request, responseStatus ResponseStatus) Response {

	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     responseStatus.StatusCode,
		StatusText: responseStatus.StatusText,
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

func GetFile(path string) Response {
	content, err := fileUtils.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return NOT_FOUND()
	}
	contentLength := fmt.Sprintf("%d", len(content))

	response := Response{
		Protocol:   "HTTP/1.1",
		Status:     200,
		StatusText: "OK",
		Body:       content,
		Headers: map[string]string{
			"Content-Type":   "application/octet-stream",
			"Content-Length": contentLength,
		},
	}

	return response
}

func UploadFile(path string, body string, responseStatus ResponseStatus) Response {
	err := fileUtils.WriteFile(path, body)
	var response Response
	if err != nil {
		fmt.Println(err)
		response = Response{
			Protocol:   "HTTP/1.1",
			Status:     BAD_REQUEST400.StatusCode,
			StatusText: BAD_REQUEST400.StatusText,
		}
	} else {
		response = Response{
			Protocol:   "HTTP/1.1",
			Status:     responseStatus.StatusCode,
			StatusText: responseStatus.StatusText,
		}
	}
	return response
}

func ParseResponse(r Response) string {

	//First line of the response
	response := fmt.Sprintf("%s %d %s\r\n", r.Protocol, r.Status, r.StatusText)

	//Headers
	for key, value := range r.Headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	//Empty line
	response += "\r\n"

	//Body
	response += r.Body

	//Empty line
	//response += "\r\n"

	return response

}
