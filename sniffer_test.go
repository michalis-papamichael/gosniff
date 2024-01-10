package gosniff

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/hashicorp/go-hclog"
)

func TestSniffPackets(t *testing.T) {
	log := hclog.Default()
	filter := "tcp"
	sniffer := Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}
	pkts, err := sniffer.StartSniff()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
	go func() {
		for p := range pkts {
			// do something
			fmt.Println(p)
		}
		log.Info("Packets channel closed")
	}()
	<-time.After(15 * time.Second)
	log.Info("Closing packet sniffer")
	stats, _ := sniffer.CloseAndGetStats(false)
	fmt.Printf("\n Stats Received: %v\n", stats.PacketsReceived)
	<-time.After(10 * time.Second)
}
