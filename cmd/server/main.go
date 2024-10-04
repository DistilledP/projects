package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/DistilledP/lungfish/internal"
	"github.com/DistilledP/lungfish/internal/libs"
	"github.com/DistilledP/lungfish/internal/parser"
	"github.com/DistilledP/lungfish/internal/types"
)

func main() {
	config := libs.GetServices().GetConfig()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PublicPort))
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	fmt.Printf("Server started on port %d\n", config.PublicPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERR accepting connection:", err)
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	// connection is opened for every command, need to catch the disconnect and handle appropriately.
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())

	bufferedConn := bufio.NewReader(conn)

	firstByte, _ := bufferedConn.Peek(1)

	var cmd types.Command
	if firstByte[0] == parser.UrpStartASCII {
		cmd = parser.ParseRedisURP(bufferedConn)
	} else {
		cmd = parser.ParseRaw(bufferedConn)
	}

	internal.DispatchCommand(conn, cmd)
}
