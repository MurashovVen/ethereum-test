SHELL := /bin/bash

gen: protoc mockgen

protoc:
	protoc --go_out=./pkg/grpc --go-grpc_out=./pkg/grpc --proto_path=./api/proto ethereum.proto

mockgen:
	mockgen --source=./internal/domain/contract.go --destination=internal/domain/mocks/ethereum_client.go --package=mocks EthereumClient

run:
	go run ./cmd/