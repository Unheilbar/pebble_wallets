test-state:
	go test -v core/state/*.go
	rm -rf test
test-vm:
	go test -v ./core/vm/*.go  
	rm -rf test