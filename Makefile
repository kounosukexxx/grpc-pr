.PHONY: protobuf
protobuf: protobuf-go protobuf-doc

.PHONY: protobuf-go
protobuf-go:
	rm -rf pb
	mkdir -p pb
	protoc --proto_path=proto/. --go-grpc_opt require_unimplemented_servers=false,paths=source_relative --go-grpc_out pb/ --go_opt paths=source_relative --go_out pb/ proto/**/*.proto

.PHONY: protobuf-doc
protobuf-doc:
	rm -rf docs/protobuf_schema
	mkdir -p docs/protobuf_schema
	protoc --doc_out=html,rest.html:docs/protobuf_schema proto/rest/*.proto


.PHONY: evans
evans:
	evans -r repl --port 8080