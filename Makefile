install_libcap:
	sudo apt-get install libpcap-dev

build_and_run:
	go build -o build/sniff sniff.go
	sudo ./build/sniff

compile_and_run_tests:
	go test -c
	mv ./sniff.go.test ./build/tests/sniff.go.test
	sudo ./build/tests/sniff.go.test