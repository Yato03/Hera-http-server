package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func handleConnection(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	size, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to recieve any information")
	}

	req := string(buffer[:size])

	request := http.ParseRequest(req)

	response := http.Controller(request)

	responseString := http.ParseResponse(response)

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(responseString); err != nil {
		fmt.Println("Unable to send data")
	}

	writer.Flush()
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	for {
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
		go handleConnection(conn)
	}

}
