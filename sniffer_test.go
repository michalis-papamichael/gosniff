package gosniff

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/gopacket/pcap"
)

func TestSniffPackets(t *testing.T) {
	// iname := "lo"
	filter := ""
	s := Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}
	pkts, err := s.StartSniff()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
	go func() {
		for p := range pkts {
			// do something
			fmt.Println(p)
		}
		fmt.Println("Packets Channel Closed")
	}()
	<-time.After(10 * time.Second)
	s.Close()
	<-time.After(10 * time.Second)
}
