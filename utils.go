package gosniff

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func PrintDeviceInterfaces() error {
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
			fmt.Printf("\t\tLoopback: %v\n", addr.IP.IsLoopback())
			fmt.Printf("\t\tSubnet mask: %s\n", addr.Netmask)
			fmt.Printf("\t\tBroadcast address: %s\n", addr.Broadaddr)
			fmt.Printf("\t\tPeer-to-Peer dest address: %s\n", addr.P2P)
		}
		fmt.Println()
	}
	return nil
}
