package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind port 6379")
		os.Exit(1)
	}

	defer l.Close()

	for {
		var conn net.Conn
		conn, err = l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+PONG\r\n"))
	}

	conn.Close()
}
