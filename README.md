# gosniff

[![Go Doc](https://godoc.org/github.com/huandu/go-clone?status.svg)](https://pkg.go.dev/github.com/michalis-papamichael/gosniff)

A simple packet sniffer that makes use of gopacket.

## Quick Start

### Installation
```makefile
go get github.com/michalis-papamichael/gosniff
```
### Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/michalis-papamichael/gosniff"
)

func main(){
	filter := "tcp"
	sniffer := gosniff.Sniffer{InterfaceName: nil, BpfFilterExpr: &filter,
		SnapshotLength: 1024, Duration: 500 * time.Microsecond, Promiscuous: false}

	pkts, err := sniffer.Start()
	if err != nil {
		panic(err)
	}

	go func() {
		for p := range pkts {
			// do something
			fmt.Println(p)
		}
	}()

 	<-time.After(15 * time.Second)
 	fmt.Println("Closing packet sniffer")
 	stats, _ := sniffer.Stop(true)
	fmt.Printf("\n Packets Received: %v\n", stats.PacketsReceived)
}
```

## License
gosniff is released under the [MIT License](https://github.com/michalis-papamichael/gosniff/blob/main/LICENSE).
