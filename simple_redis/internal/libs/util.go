package libs

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

func BytesToInteger[T constraints.Integer](val []byte) T {
	numStr := ""
	for _, v := range val {
		numStr = fmt.Sprintf("%s%c", numStr, v)
	}

	ret, _ := strconv.Atoi(numStr)

	return T(ret)
}
