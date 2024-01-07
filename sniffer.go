package gosniff

import (
	"time"

	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	InterfaceName  *string
	BpfFilterExpr  *string
	SnapshotLength int32
	Duration       time.Duration
	Promiscuous    bool
	handle         *pcap.Handle
	// PacketChannel  chan gopacket.Packet
}

func (s *Sniffer) StartSniff() (chan gopacket.Packet, error) {

	iname, err := s.getInterfaceName()
	if err != nil {
		return nil, err
	}
	handle, err := pcap.OpenLive(*iname, s.SnapshotLength, s.Promiscuous, s.Duration)
	if err != nil {
		return nil, err
	}
	s.handle = handle
	// defer handle.Close()

	bpfExpr := s.getBpfFIlterExpr()
	if bpfExpr != nil {
		if err := handle.SetBPFFilter(*bpfExpr); err != nil {
			return nil, err
		}
	}

	pktChan := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	return pktChan, nil
}

func (s *Sniffer) Close() {
	s.handle.Close()
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
