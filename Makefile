.PHONY: protobuf
protobuf: protobuf-go protobuf-doc

.PHONY: protobuf-go
protobuf-go:
	rm -rf pb
	mkdir -p pb
	protoc -I . --go-grpc_opt require_unimplemented_servers=false,paths=source_relative --go-grpc_out pb/ --go_opt paths=source_relative --go_out pb/ proto/*.proto

.PHONY: protobuf-doc
protobuf-doc:
	rm -rf docs/protobuf_schema
	mkdir -p docs/protobuf_schema
	protoc --doc_out=html,protobuf.html:docs/protobuf_schema proto/*.proto

.PHONY: firebase
firebase:
	firebase emulators:start

.PHONY: clear
clear:
	rm pubsub-debug.log
	rm firestore-debug.log
	rm ui-debug.log

.PHONY: app
app:
	docker exec -it dev-app-1 /bin/sh

.PHONY: lint
lint:
	prototool lint