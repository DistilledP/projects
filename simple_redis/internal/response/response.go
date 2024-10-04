package response

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func SimpleString(value string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", value))
}

func BulkString(value string, empty bool) []byte {
	if empty {
		return []byte("$-1\r\n")
	}

	return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value))
}

func Number[T constraints.Integer](value T) []byte {
	return []byte(fmt.Sprintf(":%d\r\n", value))
}

func Array[T string | constraints.Integer](value []T) []byte {
	out := fmt.Sprintf("*%d\r\n", len(value))

	for _, val := range value {
		valStr := fmt.Sprintf("%v", val)
		out = fmt.Sprintf("%s$%d\r\n%s\r\n", out, len(valStr), valStr)
	}

	return []byte(out)
}

func Error(err error) []byte {
	return []byte(fmt.Sprintf("-Error %s\r\n", err))
}
