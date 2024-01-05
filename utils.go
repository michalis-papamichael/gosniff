package gosniff

import (
	"net"
)

func GetPermanentMacAddress() (*net.Interface, error) {
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
