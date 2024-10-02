package parser

import (
	"bufio"
	"strings"

	"github.com/DistilledP/lungfish/internal/types"
)

func ParseRaw(reader *bufio.Reader) types.Command {
	bits, _ := reader.ReadBytes(byte(10))
	fullCmd := strings.Split(string(bits[0:len(bits)-2]), " ")

	if len(fullCmd[0]) == 0 {
		return types.Command{}
	}

	return types.Command{
		Name: strings.ToLower(fullCmd[0]),
		Args: fullCmd[1:],
	}
}
