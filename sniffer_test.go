package gosniff

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/hashicorp/go-hclog"
)

func TestSniffPackets(t *testing.T) {
	l := hclog.Default()
	// iname := "wlp2s0"
	filter := "tcp"
	s := Sniffer{InterfaceName: /*&iname*/ nil, BpfFilterExpr: &filter,
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
		l.Info("Packets channel closed")
	}()
	<-time.After(15 * time.Second)
	l.Info("Closing packet sniffer")
	stats := s.Close(true)
	fmt.Printf("\n Stats %v\n", stats)
	<-time.After(10 * time.Second)
}
