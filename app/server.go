package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server running on port 4221")

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	bufr := make([]byte, 1024)
	_, err = conn.Read(bufr)
	if err != nil {
		fmt.Println("Failed to read data from the connection")
		os.Exit(1)
	}

	reponse := []byte("HTTP/1.1 200 OK\r\n\r\n")

	n, err := conn.Write(reponse)
	if err != nil {
		fmt.Println("Failed to write data to the connection")
		os.Exit(1)
	}

	fmt.Println("Written bytes: ", n)
}
