install_libcap:
	sudo apt-get install libpcap-dev

build_and_run:
	go build sniff.go
	sudo ./sniff