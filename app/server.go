package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	OK        = "HTTP/1.1 200 OK\r\n\r\n"
	NOT_FOUND = "HTTP/1.1 404 NOT FOUND\r\n\r\n"
)

func main() {
	// Listen on port 4221
	l := try_listen(4221)
	fmt.Println("Server running on port 4221")

	// Accept connection from client
	conn := try_accept(l)
	defer conn.Close()

	// Read data from the connection
	bufr := make([]byte, 1024)
	var requestBuilder strings.Builder
	readConnectionData(conn, bufr, &requestBuilder)

	// print requestBuilder
	fmt.Println(requestBuilder.String())

	request := strings.Split(requestBuilder.String(), "\r\n")
	start_line := request[0]
	path := strings.Split(start_line, " ")[1]

	// Write data to the connection
	reponse := OK
	if path != "/" {
		reponse = NOT_FOUND
	}

	n, err := conn.Write([]byte(reponse))
	if err != nil {
		fmt.Println("Failed to write data to the connection")
		os.Exit(1)
	}

	fmt.Println("Written bytes: ", n)
}

func readConnectionData(conn net.Conn, bufr []byte, requestBuilder *strings.Builder) {
	for {
		n, err := conn.Read(bufr)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Failed to read data from the connection")
			os.Exit(1)
		}
		requestBuilder.Write(bufr[:n])
		if strings.Contains(requestBuilder.String(), "\r\n\r\n") {
			break
		}
	}
}

// Helper function to accept connection
func try_accept(l net.Listener) net.Conn {
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection", err.Error())
		os.Exit(1)
	}
	return conn
}

// Helper function to listen on port 4221
func try_listen(port int) net.Listener {
	l, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Failed to bind to port" + strconv.Itoa(port))
		os.Exit(1)
	}
	return l
}
