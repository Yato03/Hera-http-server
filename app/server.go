package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func manageConnection(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	size, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to recieve any information")
	}

	req := string(buffer[:size])

	var _, path, _ string
	lines := strings.Split(req, "\r\n")
	if len(lines) > 0 {
		firstLine := strings.Split(lines[0], " ")
		//method = firstLine[0]
		path = firstLine[1]
		//protocol = firstLine[2]
		fmt.Println(path)
		headers := make(map[string]string)
		for _, line := range lines[1:] {
			if line == "" {
				break
			}

			header := strings.SplitN(line, ":", 2)
			fmt.Println(header)
			headerName := strings.TrimSpace(header[0])
			headerValue := strings.TrimSpace(header[1])
			headers[headerName] = headerValue
		}

	}

	OK := "HTTP/1.1 200 OK\r\n\r\n"
	NOT_FOUND := "HTTP/1.1 404 Not Found\r\n\r\n"

	var response string

	if path == "/" {
		response = OK
	} else {
		response = NOT_FOUND
	}

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(response); err != nil {
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
