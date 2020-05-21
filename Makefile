PROTO_DIR := proto
GRPC_ADDRESS := 127.0.0.1:7000
DATABASE_URI := tmp/db.sqlite
DATABASE_ADAPTER := sqlite3

export GRPC_ADDRESS DATABASE_URI DATABASE_ADAPTER

storepoc: build

clean:
	rm -rf docs proto/**/*.go 2>/dev/null

cli.health:
	grpcurl -plaintext $(GRPC_ADDRESS) health.HealthService.Check

cli.echo:
	grpcurl -plaintext $(GRPC_ADDRESS) store.StoreService.Echo

cli.list:
	grpcurl -plaintext $(GRPC_ADDRESS) list

cli.find:
	grpcurl -plaintext -d '{ "name": "test" }' $(GRPC_ADDRESS) store.StoreService.Create

cli.create:
	grpcurl -plaintext -d '{ "name": "test" }' $(GRPC_ADDRESS) store.StoreService.Create

sqlite:
	sqlite3 $(DATABASE_URI) -header -column -echo 'select * from stores;'

bench:
	ghz --skipTLS -n 3000 -c 20 --insecure --call store.StoreService.Echo $(GRPC_ADDRESS)

deps:
	( cd /tmp; \
		go get \
			github.com/golang/protobuf/protoc-gen-go \
			google.golang.org/grpc \
			github.com/favadi/protoc-go-inject-tag \
			github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
	)

build: proto docs
	go build -o storepoc cmd/storepoc/main.go

test: proto
	go test core/**/*

run: proto
	go run cmd/storepoc/main.go

docs: prepare-docs docs-md docs-html

prepare-docs:
	rm -rf docs 2>/dev/null; \
		mkdir docs

docs-md:
	protoc \
		-I=$(PROTO_DIR)/v1 \
		--doc_out=./docs \
		--doc_opt=html,index.html \
		$(PROTO_DIR)/v1/*.proto

docs-html:
	protoc \
		-I=$(PROTO_DIR)/v1 \
		--doc_out=./docs \
		--doc_opt=markdown,readme.md \
		$(PROTO_DIR)/v1/*.proto

.PHONY: proto
proto: proto-v1

proto-gen-v1:
	protoc \
		-I=$(PROTO_DIR)/v1 \
		--go_out=plugins=grpc:$(PROTO_DIR)/v1 \
		$(PROTO_DIR)/v1/*.proto

proto-v1: proto-gen-v1
	for file in proto/v1/*.pb.go; do \
		protoc-go-inject-tag -input=$$file; \
	done
