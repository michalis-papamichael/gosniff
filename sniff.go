package gosniff

import (
	"fmt"

	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	defaultSnapLen = 1024
)

func PrintMachineInterfaces() error {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}
	fmt.Println("Interfaces found: ")
	for _, i := range interfaces {
		fmt.Printf("\tName: %s\n", i.Name)
		fmt.Printf("\tDesc: %s\n", i.Description)
		fmt.Printf("\tFlag: %v\n", i.Flags)
		for _, addr := range i.Addresses {
			fmt.Printf("\t\tIP Address: %s\n", addr.IP)
			fmt.Printf("\t\tSubnet mask: %s\n", addr.Netmask)
			fmt.Printf("\t\tBroadcast address: %s\n", addr.Broadaddr)
			fmt.Printf("\t\tPeer-to-Peer dest address: %s\n", addr.P2P)
		}
		fmt.Println()
	}
	return nil
}

func SniffPackets(bpfFIlter string) error {
	handle, err := pcap.OpenLive("wlp2s0", defaultSnapLen, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()
	if err := handle.SetBPFFilter(bpfFIlter); err != nil {
		return err
	}
	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		fmt.Println(pkt)
	}
	return nil
}
