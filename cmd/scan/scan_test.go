package main

import (
	"testing"

	conn "github.com/Nikkely/goports/internal/conn"
)

func BenchmarkScan_Naive(b *testing.B) {
	ports := conn.MakeWellKnownPortsList()
	b.ResetTimer()
	naiveScan(DefaultTargetHost, ports, DefaultTimeOutMillSec)
}

func BenchmarkScan_Para(b *testing.B) {
	ports := conn.MakeWellKnownPortsList()
	b.ResetTimer()
	paraScan(DefaultTargetHost, ports, DefaultTimeOutMillSec)
}
