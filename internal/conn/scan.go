package conn

import (
	"net"
	"strings"
	"time"
)

// CheckTCPOpend check tcp connection and return error message.
func CheckTCPOpened(host, port string, timeoutDur time.Duration) (msg string) {
	_, err := net.DialTimeout("tcp", makeAddr(host, port), timeoutDur)
	if err != nil {
		msg = err.Error()
	}
	return
}

func makeAddr(host, port string) string {
	return strings.Join([]string{host, port}, addrSep)
}
