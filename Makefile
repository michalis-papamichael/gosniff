.PHONY:install
install:
	sudo apt-get install libpcap-dev
	go get github.com/michalis-papamichael/gosniff

.PHONY: make_dir
make_dir:
	if [ ! -d "./build" ] ; then  mkdir ./build ; if [ ! -d "./build/tests" ] ; then  mkdir ./build/tests ; fi; fi 

.PHONY:tests
tests: make_dir	
	go test -c
	mv ./gosniff.test ./build/tests/gosniff.go.test
	sudo ./build/tests/gosniff.go.test
