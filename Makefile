generate:
	protoc -I ../proto \
	--go_out ../pkg/ --go_opt paths=source_relative \
	--go-grpc_out ../pkg/ --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ../pkg/ --grpc-gateway_opt paths=source_relative \
	--openapiv2_out ../api --openapiv2_opt logtostderr=true \
	../proto/grpcGen/config.proto
all: generate
	go build -o ../cmd/bin/client ../cmd/client/main.go
	go build -o ../cmd/bin/service ../cmd/service/main.go