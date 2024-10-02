package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"maps"
	"net"
	"slices"
	"strings"

	"github.com/DistilledP/lungfish/internal/libs"
)

var kvStore map[string][]byte = make(map[string][]byte)

func main() {
	config := libs.GetServices().GetConfig()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PublicPort))
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

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
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())

	bufferedConn := bufio.NewReader(conn)

	buff := []byte{}
	for {
		v, err := bufferedConn.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("read err:", err)
			continue
		}

		fmt.Print(string(v))

		buff = append(buff, v)

		if buff[len(buff)-1] == 10 {
			dispatchCommand(conn, buff)
			buff = []byte{}
		}
	}
}

var CRLF = []byte{13, 10}
var PINGCMD = []byte{80, 73, 78, 71}
var SET = []byte{83, 69, 84}
var GET = []byte{71, 69, 84}
var KEYS = []byte{75, 69, 89, 83}

func dispatchCommand(conn net.Conn, buff []byte) {
	cmd, args := cmdAndArgs(buff)

	switch true {
	case slices.Equal(cmd, PINGCMD):
		cmdPong(conn, args)
	case slices.Equal(cmd, SET):
		cmdSet(conn, args)
	case slices.Equal(cmd, GET):
		cmdGet(conn, args)
	case slices.Equal(cmd, KEYS):
		cmdKeys(conn)
	default:
		cmdError(conn, fmt.Sprintf("unknown command: %s", string(cmd)))
		fmt.Println(buff)
	}
}

func cmdAndArgs(raw []byte) ([]byte, [][]byte) {
	raw = bytes.TrimSpace(raw)
	buffBits := bytes.Split(raw, []byte{32})
	cmd := buffBits[0]
	args := buffBits[1:]

	return cmd, args
}

func cmdPong(conn net.Conn, args [][]byte) {
	fmt.Println(args)
	conn.Write([]byte("+PONG\r\n"))
}

func cmdSet(conn net.Conn, args [][]byte) {
	if len(args) != 2 {
		cmdError(conn, fmt.Sprintf("SET: expected 2 args, got %d", len(args)))
		return
	}

	kvStore[string(args[0])] = args[1]

	conn.Write([]byte("+OK\r\n"))
}

func cmdGet(conn net.Conn, args [][]byte) {
	if len(args) != 1 {
		cmdError(conn, fmt.Sprintf("GET: expected 1 args, got %d", len(args)))
		return
	}

	if v, found := kvStore[string(args[0])]; found {
		// this is the wrong format for Redis, colon should be CRLF - need to update client
		resp := fmt.Sprintf("$%d:%s%s", len(v), v, CRLF)
		conn.Write([]byte(resp))
		return
	}

	conn.Write([]byte("$\r\n"))
}

func cmdKeys(conn net.Conn) {
	keysStrings := []string{}
	keys := maps.Keys(kvStore)
	for k := range keys {
		keysStrings = append(keysStrings, k)
	}

	// this is the wrong format for Redis, colon should be CRLF - need to update client
	if len(keysStrings) > 0 {
		conn.Write([]byte(fmt.Sprintf("*%d:%s\r\n", len(keysStrings), strings.Join(keysStrings, ":"))))
	} else {
		conn.Write([]byte(fmt.Sprintf("*%d\r\n", len(keysStrings))))
	}
}

func cmdError(conn net.Conn, errorTxt string) {
	conn.Write([]byte(fmt.Sprintf("-Error %s\r\n", errorTxt)))
}
