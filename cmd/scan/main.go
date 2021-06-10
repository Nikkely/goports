package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	conn "github.com/Nikkely/goports/internal/conn"
)

var (
	targetHostArg     string
	timeoutMillSecOpt int
	verboseOpt        bool
	singleOpt         bool
)

func parse() {
	flag.IntVar(&timeoutMillSecOpt, "t", 100, "specify timeout in milliseconds")
	flag.BoolVar(&verboseOpt, "v", false, "show detail")
	flag.BoolVar(&singleOpt, "s", true, "scan with single thread")
	flag.Parse()
	targetHostArg = flag.Arg(0)
	if len(targetHostArg) == 0 {
		targetHostArg = "127.0.0.1"
	}
}

func main() {
	parse()
	fmt.Printf("start scanning... target: %s\n", targetHostArg)
	ports := conn.MakeWellKnownPortsList()
	scanner := paraScan
	if singleOpt {
		scanner = naiveScan
	}
	results := scanner(targetHostArg, ports, time.Duration(timeoutMillSecOpt)*time.Millisecond)
	for i, port := range ports {
		if len(results[i]) == 0 {
			fmt.Printf("port: %s is open\n", port)
		} else if verboseOpt {
			fmt.Printf("port: %s %s\n", port, results[i])
		}
	}
	fmt.Println("completed.")
}

func naiveScan(addr string, ports []string, timeout time.Duration) []string {
	resultList := make([]string, len(ports))
	for i, port := range ports {
		resultList[i] = conn.IsTCPOpened(addr, port, timeout)
	}
	return resultList
}

func paraScan(addr string, ports []string, timeout time.Duration) []string {
	resultList := make([]string, len(ports))
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
