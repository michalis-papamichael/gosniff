package gosniff

import (
	"testing"
	"time"
)

func TestSniffPackets(t *testing.T) {
	// iname := "lo"
	filter := "arp"
	s := Sniffer{InterfaceName: nil, BpfFilterExpr: &filter, SnapshotLength: 1024, Duration: 10 * time.Second}
	err := s.StartSniff()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
}
