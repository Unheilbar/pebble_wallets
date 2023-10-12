test-state:
	go test -v core/state/*.go
	rm -rf testdb
test-vm:
	go test -v ./core/vm/*.go  
	rm -rf testdb
test-processor:
	go test -v -run 'Test__StateProcessor' ./tests/*.go
	rm -rf testdb
build-contract:
	docker run -v `pwd`/contracts:/contracts/sources ethereum/solc:stable -o /contracts/sources --abi --bin /contracts/sources/contract.sol --overwrite --via-ir --optimize
	abigen --abi contracts/Storage.abi --bin contracts/Storage.bin --pkg binding --type Storage --out binding/storage.go