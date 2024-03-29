package gosniff

import (
	"net"
	"time"

	"github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	// Device interface Name
	InterfaceName *string
	// BPF expression for filtering
	BpfFilterExpr *string
	// The maximum length of the packets captured
	SnapshotLength int32
	// The duration of the sniffer were packets will be consumed
	Duration time.Duration
	// Determines whether to use Promiscuous mode
	Promiscuous bool
	// pcap Handle
	handle *pcap.Handle
}

// Starts the sniffing process.
func (s *Sniffer) Start() (chan gopacket.Packet, error) {
	iname, err := s.getInterfaceName()
	if err != nil {
		return nil, err
	}
	handle, err := pcap.OpenLive(*iname, s.SnapshotLength, s.Promiscuous, s.Duration)
	if err != nil {
		return nil, err
	}
	s.handle = handle
	if s.BpfFilterExpr != nil {
		if err := handle.SetBPFFilter(*s.BpfFilterExpr); err != nil {
			return nil, err
		}
	}
	pktChan := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	return pktChan, nil
}

// Stop's the sniffing process with option of returning stats.
func (s *Sniffer) Stop(getStats bool) (*pcap.Stats, error) {
	defer s.handle.Close()
	if getStats {
		stat, err := s.handle.Stats()
		if err != nil {
			return stat, err
		}
		return stat, nil
	}
	return nil, nil
}

func (s *Sniffer) GetHandle() *pcap.Handle {
	return s.handle
}

func (s *Sniffer) getInterfaceName() (*string, error) {
	if s.InterfaceName == nil {
		Interface, err := getPhysicalInterface()
		if err != nil {
			return nil, err
		}
		return &Interface.Name, nil
	} else {
		return s.InterfaceName, nil
	}
}

func getPhysicalInterface() (*net.Interface, error) {
	intfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range intfs {
		if len(i.HardwareAddr.String()) > 0 {
			return &i, nil
		}
	}
	return nil, nil
}
