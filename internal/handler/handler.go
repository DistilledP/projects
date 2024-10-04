package command

import (
	"fmt"
	"maps"
	"net"
)

func Del(conn net.Conn, args []string) {
	deletedCount := 0
	for _, k := range args {
		if _, found := kvStore[k]; found {
			delete(kvStore, k)
			deletedCount++
		}
	}

	conn.Write([]byte(fmt.Sprintf(":%d\r\n", deletedCount)))
}

func Error(conn net.Conn, errorTxt string) {
	conn.Write([]byte(fmt.Sprintf("-Error %s\r\n", errorTxt)))
}

func Keys(conn net.Conn, _ []string) {
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

func Ping(conn net.Conn, args []string) {
	if len(args) > 0 {
		conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0]), args[0])))
	} else {
		conn.Write([]byte("+PONG\r\n"))
	}
}
