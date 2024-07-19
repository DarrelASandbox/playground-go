PROTO_DIR = ./adapters/grpcserver

generate-proto:
	cd $(PROTO_DIR) && protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	greet.proto

b:
	go test ./...

ut:
	go test -short ./...
