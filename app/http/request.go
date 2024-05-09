package http

import (
	"strings"
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Headers  map[string]string
	Body     string
}

func ParseRequest(req string) Request {

	var request Request

	lines := strings.Split(req, "\r\n")
	if len(lines) > 0 {
		firstLine := strings.Split(lines[0], " ")
		request.Method = firstLine[0]
		request.Path = firstLine[1]
		request.Protocol = firstLine[2]
		request.Headers = make(map[string]string)

		n := 1

		for i, line := range lines[1:] {
			if line == "" {
				n = i + 1
				break
			}

			header := strings.SplitN(line, ":", 2)
			headerName := strings.TrimSpace(header[0])
			headerValue := strings.TrimSpace(header[1])
			request.Headers[headerName] = headerValue
		}
		request.Body = strings.Join(lines[n:], "")
	}
	return request
}
