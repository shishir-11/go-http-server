package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Println("Error reading request")
	}
	// log.Println(buf)
	// if strings.HasPrefix(stri)

	if req.URL.Path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else if req.URL.Path[:6] == "/echo/" {
		respString := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\n%s", req.URL.Path[6:])
		conn.Write([]byte(respString))
	} else {
		// log.Println(req.URL.Path[:7])
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	connection, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	HandleConn(connection)
}
