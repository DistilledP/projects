package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"net"
	"strings"
	"time"

	"github.com/DistilledP/lungfish/internal/libs"
	"github.com/DistilledP/lungfish/internal/parser"
	"github.com/DistilledP/lungfish/internal/types"
)

var kvStore map[string]types.Value = make(map[string]types.Value)

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

	dispatchCommand(conn, cmd)
}

var CRLF = []byte{13, 10}

func dispatchCommand(conn net.Conn, cmd types.Command) {
	switch true {
	case cmd.Name == "set":
		cmdSet(conn, cmd.Args)
	case cmd.Name == "get":
		cmdGet(conn, cmd.Args)
	case cmd.Name == "del":
		cmdDel(conn, cmd.Args)
	case cmd.Name == "keys":
		cmdKeys(conn, cmd.Args)
	case cmd.Name == "ping":
		cmdPong(conn, cmd.Args)
	default:
		cmdError(conn, fmt.Sprintf("unknown command: %s", strings.ToUpper(cmd.Name)))
	}
}

func cmdPong(conn net.Conn, args []string) {
	if len(args) > 0 {
		conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0]), args[0])))
	} else {
		conn.Write([]byte("+PONG\r\n"))
	}
}

func cmdSet(conn net.Conn, args []string) {
	if len(args) != 2 {
		cmdError(conn, fmt.Sprintf("SET: expected 2 args, got %d", len(args)))
		return
	}

	if v, found := kvStore[string(args[0])]; found {
		v.Val = []byte(args[1])
		v.DateModified = time.Now()
		kvStore[string(args[0])] = v
	} else {
		kvStore[string(args[0])] = types.Value{Val: []byte(args[1]), DateCreated: time.Now(), DateModified: time.Now()}
	}

	conn.Write([]byte("+OK\r\n"))
}

func cmdGet(conn net.Conn, args []string) {
	if len(args) != 1 {
		cmdError(conn, fmt.Sprintf("GET: expected 1 args, got %d", len(args)))
		return
	}

	if v, found := kvStore[string(args[0])]; found {
		resp := fmt.Sprintf("$%d\r\n%s%s", len(v.Val), v.Val, CRLF)
		conn.Write([]byte(resp))
		return
	}

	conn.Write([]byte("$-1\r\n"))
}

func cmdDel(conn net.Conn, args []string) {
	deletedCount := 0
	for _, k := range args {
		if _, found := kvStore[k]; found {
			delete(kvStore, k)
			deletedCount++
		}
	}

	conn.Write([]byte(fmt.Sprintf(":%d\r\n", deletedCount)))
}

func cmdKeys(conn net.Conn, _ []string) {
	// This can probably be optimised.
	keysStrings := []string{}
	keys := maps.Keys(kvStore)
	for k := range keys {
		keysStrings = append(keysStrings, k)
	}

	if len(keysStrings) > 0 {
		out := fmt.Sprintf("*%d\r\n", len(keysStrings))
		for _, k := range keysStrings {
			out = fmt.Sprintf("%s$%d\r\n%s\r\n", out, len(k), k)
		}
		conn.Write([]byte(out))
	} else {
		conn.Write([]byte(fmt.Sprintf("*%d\r\n", len(keysStrings))))
	}
}

func cmdError(conn net.Conn, errorTxt string) {
	conn.Write([]byte(fmt.Sprintf("-Error %s\r\n", errorTxt)))
}
