package command

import (
	"fmt"
	"net"
	"time"

	"github.com/DistilledP/lungfish/internal/types"
)

func Get(conn net.Conn, args []string) {
	if len(args) != 1 {
		Error(conn, fmt.Sprintf("GET: expected 1 args, got %d", len(args)))
		return
	}

	if v, found := kvStore[string(args[0])]; found {
		resp := fmt.Sprintf("$%d\r\n%s%s", len(v.Val), v.Val, CRLF)
		conn.Write([]byte(resp))
		return
	}

	conn.Write([]byte("$-1\r\n"))
}

func Set(conn net.Conn, args []string) {
	if len(args) != 2 {
		Error(conn, fmt.Sprintf("SET: expected 2 args, got %d", len(args)))
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
