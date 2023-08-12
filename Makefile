

all: clean proto backend

clean:
	rm src/protos/mancala_grpc.pb.go
	rm src/protos/mancala.pb.go
	rm src/protos/generated/mancala_grpc_pb.js
	rm src/protos/generated/mancala_pb.js
	rm src/backend/backend

proto: src/protos/mancala.proto
	cd src/protos;  \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mancala.proto; \
	grpc_tools_node_protoc --js_out=import_style=commonjs,binary:generated --grpc_out=grpc_js:generated mancala.proto

backend: src/backend/server.go
	cd src/backend; \
	go build; \
	./backend

