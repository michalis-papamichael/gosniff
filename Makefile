install_libcap:
	sudo apt-get install libpcap-dev

build_and_run:
	go build -o build/sniffer sniffer.go
	sudo ./build/sniffer

.PHONY:tests
tests:
	go test -c
	mv ./sniffer.go.test ./build/tests/sniffer.go.test
	sudo ./build/tests/sniffer.go.test