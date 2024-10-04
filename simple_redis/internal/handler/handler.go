package command

import (
	"errors"
	"net"

	"github.com/DistilledP/projects/simple_redis/internal/response"
	"github.com/DistilledP/projects/simple_redis/internal/types"
)

func Del(store types.StorageBucket) types.CommandHandler {
	return func(conn net.Conn, args []string) {
		deleted := store.Del(args...)

		conn.Write(response.Number(deleted))
	}
}

func Error(conn net.Conn, errorTxt string) {
	conn.Write(response.Error(errors.New(errorTxt)))
}

func Keys(store types.StorageBucket) types.CommandHandler {
	return func(conn net.Conn, filters []string) {
		// need to handle filters, including empty
		indexes := store.Indexes("")
		conn.Write(response.Array(indexes))
	}
}

func Ping(conn net.Conn, args []string) {
	if len(args) > 0 {
		conn.Write(response.BulkString(args[0], false))
	} else {
		conn.Write(response.SimpleString("PONG"))
	}
}
