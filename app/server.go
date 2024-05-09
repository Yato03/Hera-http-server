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
}

func OK() Response {
	return Response{
		Protocol:   "HTTP/1.1",
		Status:     200,
		StatusText: "OK",
	}
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

	response := fmt.Sprintf("%s %d %s\r\n\r\n", r.Protocol, r.Status, r.StatusText)

	return response

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

	var response Response

	if request.Path == "/" {
		response = OK()
	} else {
		response = NOT_FOUND()
	}

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
