package main

import (
	"fmt"
	"net"
  "os"
  "bufio"
)

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

  http_pdu := "HTTP/1.1 200 OK\r\n\r\n"
  reader := bufio.NewReader(conn)
  _, err = reader.ReadString('\n')
  if err != nil {
    fmt.Println("Failed to recieve any information")
  }

  writer := bufio.NewWriter(conn)
  if _,err := writer.WriteString(http_pdu); err != nil {
    fmt.Println("Unable to send data")
  }

  writer.Flush()
  
}
