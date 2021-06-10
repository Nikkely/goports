package conn

import (
	"net"
	"strings"
	"time"
)

// IsTCPOpend check tcp connection with host, port
func IsTCPOpened(host, port string, timeoutDur time.Duration) string {
	_, err := net.DialTimeout("tcp", makeAddr(host, port), timeoutDur)
	return err.Error()
}

func makeAddr(host, port string) string {
	return strings.Join([]string{host, port}, addrSep)
}
