PROTO_DIR := proto
GRPC_ADDRESS := 127.0.0.1:7000
DATABASE_URI := db.sqlite
DATABASE_ADAPTER := sqlite3

export GRPC_ADDRESS DATABASE_URI DATABASE_ADAPTER

storepoc: build

cli.health:
	grpcurl -plaintext -d '{}' $(GRPC_ADDRESS) health.HealthService.Check

cli.echo:
	grpcurl -plaintext $(GRPC_ADDRESS) store.StoreService.Echo

cli.list:
	grpcurl -plaintext $(GRPC_ADDRESS) list

cli.create:
	grpcurl -plaintext -d '{ "name": "test", "uri": "https://urisample.com" }' $(GRPC_ADDRESS) store.StoreService.Create

sqlite:
	sqlite3 $(DATABASE_URI) -header -column -echo 'select * from stores;'

bench:
	ghz --skipTLS -n 3000 -c 20 --insecure --call store.StoreService.Echo $(GRPC_ADDRESS)

deps:
	( cd /tmp; \
		go get -u github.com/golang/protobuf/protoc-gen-go; \
		go get -u google.golang.org/grpc; \
	)

.PHONY: proto
proto:
	protoc \
		-I=$(PROTO_DIR) \
		--go_out=plugins=grpc:$(PROTO_DIR) \
		$(PROTO_DIR)/*.proto

build: proto
	go build -o storepoc cmd/storepoc/main.go

test: proto
	go test core/**/*

run: proto
	go run cmd/storepoc/main.go
