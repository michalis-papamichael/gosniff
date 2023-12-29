package main

import (
	"fmt"

	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	defaultSnapLen = 262144
)

func main() {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	fmt.Println("Interfaces found: ")
	for _, i := range interfaces {
		fmt.Printf("\tName: %s\n", i.Name)
		fmt.Printf("\tDesc: %s\n", i.Description)
		fmt.Printf("\tFlag: %v\n", i.Flags)
		fmt.Println()
	}
	// handle, err := pcap.OpenLive("eth0", defaultSnapLen, true, pcap.BlockForever)
	// if err != nil {
	// 	panic(err)
	// }
	// defer handle.Close()
	// if err := handle.SetBPFFilter("port 3030"); err != nil {
	// 	panic(err)
	// }
	// packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	// for pkt := range packets {
	// 	fmt.Println(pkt.Metadata().AncillaryData...)
	// }
}
