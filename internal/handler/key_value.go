package command

import (
	"fmt"
	"net"

	"github.com/DistilledP/lungfish/internal/types"
)

func Get(store types.StorageBucket) types.CommandHandler {
	return func(conn net.Conn, args []string) {
		if len(args) != 1 {
			Error(conn, fmt.Sprintf("GET: expected 1 args, got %d", len(args)))
			return
		}

		key := string(args[0])
		matches := store.Find(key)
		if len(matches) > 0 {
			match := matches[key]
			resp := fmt.Sprintf("$%d\r\n%s%s", len(match.Val), match.Val, CRLF)
			conn.Write([]byte(resp))
			return
		}

		conn.Write([]byte("$-1\r\n"))
	}
}

func Set(store types.StorageBucket) types.CommandHandler {
	return func(conn net.Conn, args []string) {
		if len(args) != 2 {
			Error(conn, fmt.Sprintf("SET: expected 2 args, got %d", len(args)))
			return
		}

		key := string(args[0])
		value := string(args[1])

		store.Add(key, value)

		conn.Write([]byte("+OK\r\n"))
	}
}
