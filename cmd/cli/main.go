package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", fmt.Sprintf("localhost:%d", 6379))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Write([]byte("Hello from cli"))
	conn.(*net.TCPConn).CloseWrite()

	resp, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(resp))
}
