package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	conn "github.com/Nikkely/goports/internal/conn"
)

const (
	ArgOrderOfTargetHost  = 0
	DefaultTargetHost     = "127.0.0.1"
	DefaultTimeOutMillSec = 100
)

func main() {
	timeoutMillSec := flag.Int("timeout", DefaultTimeOutMillSec, "specify timeout in milliseconds")
	flag.Parse()
	targetHost := flag.Arg(ArgOrderOfTargetHost)
	if len(targetHost) == 0 {
		targetHost = DefaultTargetHost
	}

	fmt.Printf("start scanning... target: %s\n", targetHost)
	ports := conn.MakeWellKnownPortsList()
	results := naiveScan(targetHost, ports, time.Duration(*timeoutMillSec)*time.Millisecond)
	for i, port := range ports {
		if results[i] {
			fmt.Printf("port: %s is open\n", port)
		}
	}
	fmt.Println("completed.")
}

func naiveScan(addr string, ports []string, timeout time.Duration) []bool {
	resultList := make([]bool, len(ports))
	for i, port := range ports {
		resultList[i] = conn.IsTCPOpened(addr, port, timeout)
	}
	return resultList
}

func paraScan(addr string, ports []string, timeout time.Duration) []bool {
	resultList := make([]bool, len(ports))
	var wg sync.WaitGroup
	for i := range ports {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			resultList[x] = conn.IsTCPOpened(addr, ports[x], timeout)
		}(i)
	}
	wg.Wait()
	return resultList
}
