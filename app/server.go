package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/fileUtils"
	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func handleConnection(conn net.Conn, router *http.Router) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	size, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to recieve any information")
	}

	req := string(buffer[:size])

	request := http.ParseRequest(req)

	response := router.Handle(request)

	responseString := http.ParseResponse(response)

	fmt.Println(responseString)

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(responseString); err != nil {
		fmt.Println("Unable to send data")
	}

	writer.Flush()
}

func main() {

	// Arguments
	var (
		flDirectory = flag.String("directory", ".", "Directory to serve files from")
	)

	flag.Parse()

	fmt.Println("Serving files from", *flDirectory)

	if *flDirectory != "" {
		fileUtils.MakeConfigurationFile(*flDirectory)
	} else {
		fileUtils.CleanConfiguration()
	}

	// Uncomment this block to pass the first stage¡
	l, err := net.Listen("tcp", "0.0.0.0:4221")

	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Listening on 0.0.0.0:4221")

	router := &http.Router{}

	router.AddHandler(&http.HomeController{Path: "/", Method: "GET"})
	router.AddHandler(&http.EchoController{Path: "/echo/", Method: "GET"})
	router.AddHandler(&http.UserAgentController{Path: "/user-agent", Method: "GET"})
	router.AddHandler(&http.GetFilesController{Path: "/files", Method: "GET"})
	router.AddHandler(&http.UploadFileController{Path: "/files", Method: "POST"})

	var conn net.Conn

	for {
		conn, err = l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		//Accepted connection
		go handleConnection(conn, router)
	}

}
