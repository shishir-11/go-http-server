package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func HandleConn(conn net.Conn, dir string) {
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
		customHeaders := ""
		// if v,ok := req.Header["Accept-Encoding"];ok{

		// }
		// log.Println(req.Header["Accept-Encoding"])
		// log.Println(req.Header["Accept-Encoding"])
		if req.Header["Accept-Encoding"] != nil && req.Header["Accept-Encoding"][0] == "gzip" {
			customHeaders += "Content-Encoding: gzip\r\n"
		}
		respString := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n%s\r\n%s", len(req.URL.Path[6:]), customHeaders, req.URL.Path[6:])
		conn.Write([]byte(respString))
	} else if len(req.URL.Path) >= 11 && req.URL.Path[:11] == "/user-agent" {
		cont := req.Header["User-Agent"]
		log.Println(cont[0])
		respString := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(cont[0]), cont[0])
		conn.Write([]byte(respString))

	} else if strings.HasPrefix(req.URL.Path, "/files/") {

		path := req.URL.Path
		if strings.HasPrefix(path, "/files/") {

			directory := os.Args[2]

			fileName := strings.TrimPrefix(path, "/files/")

			if req.Method == "POST" {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					log.Printf("Unable to read body: %v", err)
					conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
					return
				}
				// log.Println(body)
				err = os.WriteFile(dir+fileName, body, 0644)
				if err != nil {
					log.Printf("Unable to write file: %v", err)
					conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
				}
				conn.Write([]byte("HTTP/1.1 201 Created\r\n\r\n"))
			} else {
				data, err := os.ReadFile(directory + fileName)

				if err != nil {

					conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))

				} else {

					conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: " + strconv.Itoa(len(data)) + "\r\n\r\n" + string(data) + "\r\n\r\n"))

				}

			}

		}

	} else {
		// log.Println(req.URL.Path[:7])
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	dir := flag.String("directory", "", "")
	flag.Parse()

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {
		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleConn(connection, *dir)
	}

}
