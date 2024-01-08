package gosniff

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/gopacket/pcap"
)

func TestUtils(t *testing.T) {
	i, err := GetPhysicalInterface()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
	fmt.Printf("\nInterface %v\n\n", i)
	err = PrintDeviceInterfaces()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
}

func TestSniffPackets(t *testing.T) {
	iname := "wlp2s0"
	filter := ""
	s := Sniffer{InterfaceName: &iname, BpfFilterExpr: &filter,
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
