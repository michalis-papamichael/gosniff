# go-sniff

A simple packet sniffer that makes use of gopacket.

## Quick Start

### Installation
```makefile
make install
```
### Usage

```go
package example

import (
	gosniff "github.com/michalis-papamichael/go-sniff"
)

func main(){
	ffilter := "tcp"
	sniffer := gosniff.Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}

	pkts, err := sniffer.StartSniff()
	if err != nil {
		t.Fatal(err)
		panic(err)
	}

	go func() {
		for p := range pkts {
		// do something
		}
	}()

 	<-time.After(15 * time.Second)
 	fmt.Println("Closing packet sniffer")
 	stats, _ := sniffer.Close(true)
}
```

## License
go-sniff is released under the [MIT License](https://github.com/michalis-papamichael/go-sniff/blob/main/LICENSE).
