package main

import (
	"testing"
	"time"

	conn "github.com/Nikkely/goports/internal/conn"
)

func BenchmarkScan_Naive(b *testing.B) {
	parse()
	ports := conn.MakeWellKnownPortsList()
	b.ResetTimer()
	naiveScan(targetHostArg, ports, time.Duration(timeoutMillSecOpt)*time.Millisecond)
}

func BenchmarkScan_Para(b *testing.B) {
	parse()
	ports := conn.MakeWellKnownPortsList()
	b.ResetTimer()
	paraScan(targetHostArg, ports, time.Duration(timeoutMillSecOpt)*time.Millisecond)
}
