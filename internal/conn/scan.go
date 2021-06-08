package conn

import (
	"net"
	"strings"
	"time"
)

// IsTCPOpend check tcp connection with host, port
func IsTCPOpened(host, port string, timeoutDur time.Duration) (ok bool) {
	_, err := net.DialTimeout("tcp", makeAddr(host, port), timeoutDur)
	if err == nil {
		ok = true
	}
	return
}

func makeAddr(host, port string) string {
	return strings.Join([]string{host, port}, addrSep)
}
