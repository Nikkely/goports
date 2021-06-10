package conn

import (
	"net"
	"strings"
	"time"
)

// CheckTCPOpend check tcp connection and return error message.
func CheckTCPOpened(host, port string, timeoutDur time.Duration) string {
	_, err := net.DialTimeout("tcp", makeAddr(host, port), timeoutDur)
	return err.Error()
}

func makeAddr(host, port string) string {
	return strings.Join([]string{host, port}, addrSep)
}
