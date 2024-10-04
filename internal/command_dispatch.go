package internal

import (
	"fmt"
	"net"
	"strings"

	hdlr "github.com/DistilledP/lungfish/internal/handler"
	"github.com/DistilledP/lungfish/internal/types"
)

var commandMap map[string]types.CommandHandler = map[string]func(net.Conn, []string){
	hdlr.CmdDel:  hdlr.Del,
	hdlr.CmdGet:  hdlr.Get,
	hdlr.CmdKeys: hdlr.Keys,
	hdlr.CmdPing: hdlr.Ping,
	hdlr.CmdSet:  hdlr.Set,
}

func DispatchCommand(conn net.Conn, cmd types.Command) {
	cmdFunc, found := commandMap[cmd.Name]
	if !found {
		hdlr.Error(conn, fmt.Sprintf("unknown command: %s", strings.ToUpper(cmd.Name)))
		return
	}

	cmdFunc(conn, cmd.Args)
}
