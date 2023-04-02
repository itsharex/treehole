grpc_gen:
	protoc --proto_path=../../proto \
  --go_out=rpc --go_opt=paths=source_relative \
  --go-grpc_out=rpc \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  topic.proto

grpc_gateway:
	protoc --proto_path=../../proto \
  --grpc-gateway_out logtostderr=true:rpc \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
  topic.proto

package: grpc_gen
	go build -o account