generate:
	protoc -I proto \
	--go_out . --go_opt module=config-service \
	--go-grpc_out . --go-grpc_opt module=config-service \
	--grpc-gateway_out . --grpc-gateway_opt module=config-service \
	--openapiv2_out api \
	--openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=config \
	proto/grpcGen/*.proto

all: generate
	go build -o cmd/bin/client cmd/client/main.go
	go build -o cmd/bin/service cmd/service/main.go