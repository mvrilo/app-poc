PROTO_DIR := proto
GRPC_ADDRESS := 127.0.0.1:7000
HTTP_ADDRESS := 127.0.0.1:8000
DATABASE_URI := tmp/db.sqlite
DATABASE_ADAPTER := sqlite3

export GRPC_ADDRESS HTTP_ADDRESS DATABASE_URI DATABASE_ADAPTER

storepoc: build

build-run: storepoc
	./storepoc

clean:
	rm -rf docs/* proto/**/*.go 2>/dev/null

cli.health:
	grpcurl -plaintext $(GRPC_ADDRESS) health.v1.HealthService.Check

cli.echo:
	grpcurl -plaintext $(GRPC_ADDRESS) store.v1.StoreService.Echo

cli.list:
	grpcurl -plaintext $(GRPC_ADDRESS) list

cli.find:
	grpcurl -plaintext -d '{ "name": "test" }' $(GRPC_ADDRESS) store.v1.StoreService.Create

cli.create:
	grpcurl -plaintext -d '{ "name": "" }' $(GRPC_ADDRESS) store.v1.StoreService.Create

sqlite:
	sqlite3 $(DATABASE_URI) -header -column -echo 'select * from stores;'

bench:
	ghz --skipTLS -n 3000 -c 20 --insecure --call store.StoreService.Echo $(GRPC_ADDRESS)

deps:
	( cd /tmp; \
		go get -v \
			github.com/golang/protobuf/protoc-gen-go \
			google.golang.org/grpc \
			github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
			github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
			github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
			github.com/ahmetb/govvv \
	)

build: proto docs
	govvv build -o storepoc cmd/storepoc/main.go

test: proto
	go test core/**/*

run: proto docs
	go run cmd/storepoc/main.go

prepare-docs:
	rm -rf docs/* &>/dev/null

docs-md:
	protoc \
		-I=$(PROTO_DIR)/v1 \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--doc_out=./docs \
		--doc_opt=html,index.html \
		$(PROTO_DIR)/v1/*.proto

docs-html:
	protoc \
		-I=$(PROTO_DIR)/v1 \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--doc_out=./docs \
		--doc_opt=markdown,readme.md \
		$(PROTO_DIR)/v1/*.proto

docs: prepare-docs docs-md docs-html

proto-v1:
	protoc \
		-I=$(PROTO_DIR)/v1 \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:$(PROTO_DIR)/v1 \
		--grpc-gateway_out=logtostderr=true:$(PROTO_DIR)/v1 \
		--swagger_out=logtostderr=true,allow_merge=true,merge_file_name=api:docs \
		$(PROTO_DIR)/v1/*.proto

.PHONY: proto
proto: proto-v1
