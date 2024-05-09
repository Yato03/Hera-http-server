package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Headers  map[string]string
	Body     string
}

type Response struct {
	Protocol   string
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
}

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

func parseRequest(req string) Request {

	var request Request

	lines := strings.Split(req, "\r\n")
	if len(lines) > 0 {
		firstLine := strings.Split(lines[0], " ")
		//method = firstLine[0]
		request.Path = firstLine[1]
		//protocol = firstLine[2]
		request.Headers = make(map[string]string)
		for _, line := range lines[1:] {
			if line == "" {
				break
			}

			header := strings.SplitN(line, ":", 2)
			headerName := strings.TrimSpace(header[0])
			headerValue := strings.TrimSpace(header[1])
			request.Headers[headerName] = headerValue
		}
	}
	return request
}

func parseResponse(r Response) string {

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
	response += "\r\n"

	return response

}

func controller(request Request) Response {
	var response Response

	directories := strings.Split(request.Path, "/")

	response = root(directories[1:])

	return response
}

func root(directories []string) Response {
	if len(directories) == 1 && directories[0] == "" {
		return OK("")
	} else if len(directories) >= 1 {
		if directories[0] == "echo" {
			return echo(directories[1:])
		}
	}
	return NOT_FOUND()

}

func echo(directories []string) Response {
	if len(directories) == 0 {
		return OK("ECHO!!")
	} else if len(directories) >= 1 {
		return OK(directories[0])
	}
	return NOT_FOUND()
}

func manageConnection(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	size, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to recieve any information")
	}

	req := string(buffer[:size])

	request := parseRequest(req)

	response := controller(request)

	responseString := parseResponse(response)

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(responseString); err != nil {
		fmt.Println("Unable to send data")
	}

	writer.Flush()
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage¡
	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	var conn net.Conn
	conn, err = l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	//Accepted connection
	manageConnection(conn)

}
