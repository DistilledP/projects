package command

import (
	"github.com/DistilledP/lungfish/internal/types"
)

const (
	CmdDel  = "del"
	CmdGet  = "get"
	CmdKeys = "keys"
	CmdPing = "ping"
	CmdSet  = "set"
)

var kvStore map[string]types.Value = make(map[string]types.Value)
var CRLF = []byte{13, 10}
