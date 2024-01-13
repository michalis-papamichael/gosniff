# go-sniff

[![Go Doc](https://godoc.org/github.com/huandu/go-clone?status.svg)]([https://pkg.go.dev/github.com/huandu/go-clone](https://pkg.go.dev/github.com/michalis-papamichael/gosniff))

A simple packet sniffer that makes use of gopacket.

## Quick Start

### Installation
```makefile
go get github.com/michalis-papamichael/gosniff
```
### Usage

```go
package example

import (
	gosniff "github.com/michalis-papamichael/gosniff"
)

func main(){
	ffilter := "tcp"
	sniffer := gosniff.Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}

	pkts, err := sniffer.Start()
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
 	stats, _ := sniffer.Stop(true)
}
```

## License
gosniff is released under the [MIT License](https://github.com/michalis-papamichael/gosniff/blob/main/LICENSE).
