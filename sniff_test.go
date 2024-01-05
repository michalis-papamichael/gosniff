package gosniff

import "testing"

func TestSniffPackets(t *testing.T) {
	err := SniffPackets("tcp && ip")
	if err != nil {
		t.Fatal(err)
		panic(err)
	}
}
