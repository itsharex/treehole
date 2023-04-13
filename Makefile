pkg=account

grpc_gen_account:
	protoc --proto_path=proto \
  --go_out=cmd/$(pkg)/rpc --go_opt=paths=source_relative \
  --go-grpc_out=cmd/$(pkg)/rpc \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  $(pkg).proto

grpc_gen_topic:
	protoc --proto_path=proto \
  --go_out=cmd/topic/rpc --go_opt=paths=source_relative \
  --go-grpc_out=cmd/topic/rpc \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  topic.proto

grpc_gen_code:
	protoc --proto_path=proto \
  --go_out=pkg/rpcs --go_opt=paths=source_relative \
  --go-grpc_out=pkg/rpcs \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  code.proto

#grpc_gateway:
#	protoc --proto_path=../../proto \
#  --grpc-gateway_out logtostderr=true:rpc \
#  --grpc-gateway_opt paths=source_relative \
#  --grpc-gateway_opt generate_unbound_methods=true \
#  topic.proto

#package: grpc_gen
#	go build -o account