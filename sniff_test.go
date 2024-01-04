package gosniff

import "testing"

func TestSniffPackets(t *testing.T) {
	err := PrintDeviceInterfaces()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
}
