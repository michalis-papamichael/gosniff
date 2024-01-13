package gosniff

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/gopacket/pcap"
)

func TestSniffPackets(t *testing.T) {
	filter := "tcp"
	sniffer := Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}
	pkts, err := sniffer.Start()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
	go func() {
		for p := range pkts {
			// do something
			fmt.Println(p)
		}
		fmt.Println("Packets channel closed")
	}()
	<-time.After(15 * time.Second)
	fmt.Println("Closing packet sniffer")
	stats, _ := sniffer.Stop(true)
	fmt.Printf("\n Packets Received: %v\n", stats.PacketsReceived)
	<-time.After(10 * time.Second)
}
