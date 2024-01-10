# go-sniff

A simple packet sniffer that makes use of gopacket.

## Quick Start

### Installation
```bash
go get github.com/michalis-papamichael/go-sniff
```
### Usage
This is an example of use:

```bash
sniffer := Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}
pkts, err := sniffer.StartSniff()
if err != nil {
  t.Fatal(err)
  panic(err)
}
```

## License
go-sniff is released under the [MIT License](https://github.com/michalis-papamichael/go-sniff/blob/main/LICENSE).
