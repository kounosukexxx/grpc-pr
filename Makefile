PHONY: proto
proto: 
	protoc --proto_path=proto/. --go-grpc_opt require_unimplemented_servers=false,paths=source_relative --go-grpc_out pb/ --go_opt paths=source_relative --go_out pb/ proto/*.proto

