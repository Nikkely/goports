package conn

import (
	"strconv"
)

const (
	addrSep = `:`
)

// MakeWellKnownPortsList gives well-known ports list
func MakeWellKnownPortsList() []string {
	const num = 1024
	ret := make([]string, num)
	for i := 0; i < num; i++ {
		ret[i] = strconv.Itoa(i)
	}
	return ret
}
