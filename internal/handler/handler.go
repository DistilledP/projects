package command

import (
	"fmt"
	"net"

	"github.com/DistilledP/lungfish/internal/types"
)

func Del(store types.StorageBucket) types.CommandHandler {
	return func(conn net.Conn, args []string) {
		deleted := store.Del(args...)
		conn.Write([]byte(fmt.Sprintf(":%d\r\n", deleted)))
	}
}

func Error(conn net.Conn, errorTxt string) {
	conn.Write([]byte(fmt.Sprintf("-Error %s\r\n", errorTxt)))
}

func Keys(store types.StorageBucket) types.CommandHandler {
	return func(conn net.Conn, filters []string) {
		// need to handle filters, including empty
		indexes := store.Indexes("")

		if len(indexes) > 0 {
			out := fmt.Sprintf("*%d\r\n", len(indexes))
			for _, k := range indexes {
				out = fmt.Sprintf("%s$%d\r\n%s\r\n", out, len(k), k)
			}
			conn.Write([]byte(out))
		} else {
			conn.Write([]byte("*0\r\n"))
		}
	}
}

func Ping(conn net.Conn, args []string) {
	if len(args) > 0 {
		conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(args[0]), args[0])))
	} else {
		conn.Write([]byte("+PONG\r\n"))
	}
}
