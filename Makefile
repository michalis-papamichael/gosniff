install_libcap:
	sudo apt-get install libpcap-dev

build_and_run:
	go build -o build/sniff sniff.go
	sudo ./build/sniff