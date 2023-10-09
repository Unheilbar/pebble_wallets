test-state:
	go test -v core/state/*.go
	rm -rf test
test-vm:
	go test -v ./core/vm/*.go  
	rm -rf test
build-contract:
	docker run -v `pwd`/contracts:/contracts/sources ethereum/solc:stable -o /contracts/sources --abi --bin /contracts/sources/contract.sol --overwrite 
	abigen --abi contracts/Storage.abi --bin contracts/Storage.bin --pkg binding --type Storage --out binding/storage.go