grpc_gen:
	protoc --proto_path=../../proto \
  --go_out=rpc --go_opt=paths=source_relative \
  --go-grpc_out=rpc \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  account.proto

package: grpc_gen
	go build -o account