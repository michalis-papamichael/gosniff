package gosniff

import (
	"fmt"

	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	InterfaceName  *string
	BpfFilterExpr  *string
	SnapshotLength int32
}

func (s *Sniffer) StartSniff() error {

	iname, err := s.getInterfaceName()
	if err != nil {
		return err
	}

	handle, err := pcap.OpenLive(*iname, s.SnapshotLength, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	bpfExpr := s.getBpfFIlterExpr()
	if bpfExpr != nil {
		if err := handle.SetBPFFilter(*bpfExpr); err != nil {
			return err
		}
	}
	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	for pkt := range packets {
		fmt.Println(pkt)
	}
	return nil
}

func (s *Sniffer) getInterfaceName() (*string, error) {
	if s.InterfaceName == nil {
		Interface, err := GetPermanentMacAddress()
		if err != nil {
			return nil, err
		}
		return &Interface.Name, nil
	} else {
		return s.InterfaceName, nil
	}
}

func (s *Sniffer) getBpfFIlterExpr() *string {
	return s.BpfFilterExpr
}
