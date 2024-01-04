package gosniff

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

const (
	defaultSnapLen = 1024
)

func GetCurrentUsableDeviceInterface() error {
	_, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}
	// for _, i := range interfaces {
	// }
	return nil
}

func PrintDeviceInterfaces() error {
	intfs, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, i := range intfs {
		fmt.Printf("\tInterface Name: %s\n", i.Name)
		fmt.Printf("\tMAC Addr: %s\n", i.HardwareAddr.String())
		fmt.Println()
	}
	// interfaces, err := pcap.FindAllDevs()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("Interfaces found: ")
	// for _, i := range interfaces {
	// 	fmt.Printf("\tName: %s\n", i.Name)
	// 	fmt.Printf("\tDesc: %s\n", i.Description)
	// 	fmt.Printf("\tFlag: %v\n", i.Flags)
	// 	for _, addr := range i.Addresses {
	// 		fmt.Printf("\t\tIP Address: %s\n", addr.IP)
	// 		fmt.Printf("\t\tLoopback: %v\n", addr.IP.IsLoopback())
	// 		fmt.Printf("\t\tSubnet mask: %s\n", addr.Netmask)
	// 		fmt.Printf("\t\tBroadcast address: %s\n", addr.Broadaddr)
	// 		fmt.Printf("\t\tPeer-to-Peer dest address: %s\n", addr.P2P)
	// 	}
	// 	fmt.Println()
	// }
	return nil
}

func SniffPackets(interfaceName string, bpfExpr string) error {
	handle, err := pcap.OpenLive(interfaceName, defaultSnapLen, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()
	if err := handle.SetBPFFilter(bpfExpr); err != nil {
		return err
	}
	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		fmt.Println(pkt.TransportLayer())
	}
	return nil
}

func getMacAddress(interfaceName string) {

}
