package types

import "net"

type Command struct {
	Name string
	Args []string
}

type CommandHandler = func(net.Conn, []string)
