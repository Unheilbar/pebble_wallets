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
build-complex:
#  transfer
	docker run -v `pwd`/contracts:/contracts/sources ethereum/solc:stable -o /contracts/sources --abi --bin /contracts/sources/complex/transfer.sol --overwrite --via-ir --optimize
	abigen --abi contracts/Transfer.abi --bin contracts/Transfer.bin --pkg binding --type Transfer --out binding/transfer.go
# state
	docker run -v `pwd`/contracts:/contracts/sources ethereum/solc:stable -o /contracts/sources --abi --bin /contracts/sources/complex/state.sol --overwrite --via-ir --optimize
	abigen --abi contracts/State.abi --bin contracts/State.bin --pkg binding --type State --out binding/state.go
# proxy
	docker run -v `pwd`/contracts:/contracts/sources ethereum/solc:stable -o /contracts/sources --abi --bin /contracts/sources/complex/proxy.sol --overwrite --via-ir --optimize
	abigen --abi contracts/Proxy.abi --bin contracts/Proxy.bin --pkg binding --type Proxy --out binding/proxy.go
# events
	docker run -v `pwd`/contracts:/contracts/sources ethereum/solc:stable -o /contracts/sources --abi --bin /contracts/sources/complex/events.sol --overwrite --via-ir --optimize
	abigen --abi contracts/Event.abi --bin contracts/Event.bin --pkg binding --type Events --out binding/events.go
up-full-node:
	rm -rf logs/*
	go run cmd/geth/*.go
stress-minter:
	rm -rf chaindb
	rm -rf logs/*
	# rm mem.prof
	go test -v -cpuprofile cpu.prof -run 'Test__RunStressMinter' tests/*.go -timeout 99999s