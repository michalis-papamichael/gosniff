# go-sniff

A simple packet sniffer that makes use of gopacket.

## Quick Start

### Installation
```bash
make install
```
### Usage

```bash
import gosniff "github.com/michalis-papamichael/go-sniff"

func main(){
	filter="tcp"
	sniffer := gosniffSniffer{InterfaceName: nil, BpfFilterExpr: &filter,
			SnapshotLength: 1024, Duration: pcap.BlockForever, Promiscuous: false}
	pkts, err := sniffer.StartSniff()
	if err != nil {
	  t.Fatal(err)
	  panic(err)
	}

// Do somthing with the pkts channel
	go func() {
		for p := range pkts {
			// ...
		}
	}()
	<-time.After(15 * time.Second)
	fmt.Println("Closing packet sniffer")
// -----
	stats, _ := sniffer.CloseAndGetStats(false)
}
```

## License
go-sniff is released under the [MIT License](https://github.com/michalis-papamichael/go-sniff/blob/main/LICENSE).
