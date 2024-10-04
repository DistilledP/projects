package internal

import (
	"fmt"
	"net"
	"strings"

	hdlr "github.com/DistilledP/projects/simple_redis/internal/handler"
	"github.com/DistilledP/projects/simple_redis/internal/types"
)

var commandMap map[string]types.CommandHandler

func Setup(store types.StorageBucket) {
	commandMap = map[string]types.CommandHandler{
		hdlr.CmdDel:  hdlr.Del(store),
		hdlr.CmdGet:  hdlr.Get(store),
		hdlr.CmdKeys: hdlr.Keys(store),
		hdlr.CmdPing: hdlr.Ping,
		hdlr.CmdSet:  hdlr.Set(store),
	}
}

func DispatchCommand(conn net.Conn, cmd types.Command) {
	cmdFunc, found := commandMap[cmd.Name]
	if !found {
		hdlr.Error(conn, fmt.Sprintf("unknown command: %s", strings.ToUpper(cmd.Name)))
		return
	}

	cmdFunc(conn, cmd.Args)
}
