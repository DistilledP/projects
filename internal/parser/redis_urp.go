package parser

import (
	"bufio"
	"strings"

	"github.com/DistilledP/lungfish/internal/libs"
	"github.com/DistilledP/lungfish/internal/types"
)

const (
	urpStartASCII    = 42
	urpFieldLenASCII = 36
)

func ParseRedisURP(reader *bufio.Reader) types.Command {
	header, _ := reader.ReadBytes(byte(10))
	if header[0] != urpStartASCII {
		return types.Command{}
	}

	command := []string{}
	dataLen := 0

	segments := libs.BytesToInteger[uint](header[1 : len(header)-2])
	for seg := 0; seg < int(segments*2); seg++ {
		segment, _ := reader.ReadBytes(byte(10))
		if segment[0] == urpFieldLenASCII { // header segment
			dataLen = libs.BytesToInteger[int](segment[1 : len(segment)-2])
		} else { // data segment
			command = append(command, string(segment[0:dataLen]))
		}
	}

	return types.Command{
		Name: strings.ToLower(command[0]),
		Args: command[1:],
	}
}
