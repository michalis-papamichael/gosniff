package gosniff

import "testing"

func TestSniffPackets(t *testing.T) {
	// iname := "lo"
	filter := "arp"
	s := Sniffer{InterfaceName: nil, BpfFilterExpr: &filter}
	err := s.StartSniff()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
}
