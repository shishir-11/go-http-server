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
	} else if len(req.URL.Path) > 6 && req.URL.Path[:6] == "/echo/" {
		respString := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(req.URL.Path[6:]), req.URL.Path[6:])
		conn.Write([]byte(respString))
	} else if len(req.URL.Path) >= 11 && req.URL.Path[:11] == "/user-agent" {
		cont := req.Header["User-Agent"]
		// log.Printf("%T", cont)
		// log.Println(len(cont))
		// log.Println(strings.Split(cont, "["))
		log.Println(cont[0])
		respString := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(cont[0]), cont[0])
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
	go HandleConn(connection)
}
