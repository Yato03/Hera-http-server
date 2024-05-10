package http

import "fmt"

type Response struct {
	Protocol   string
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
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
