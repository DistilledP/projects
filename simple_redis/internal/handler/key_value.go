package command

import (
	"fmt"
	"net"

	"github.com/DistilledP/projects/simple_redis/internal/response"
	"github.com/DistilledP/projects/simple_redis/internal/types"
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
			conn.Write(response.BulkString(string(matches[key].Val), false))
			return
		}

		conn.Write(response.BulkString("", true))
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

		conn.Write(response.SimpleString("OK"))
	}
}
