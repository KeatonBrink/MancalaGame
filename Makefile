

all: clean proto backend

clean:
	cd src/protos
	rm mancala_grpc.pb.go
	rm mancala.pb.go
	cd generated
	rm mancala_grpc_pb.js
	rm mancala_pb.js
	cd ../../backend
	rm backend

proto: src/protos/mancala.proto src/protos/mancala.pb.go src/protos/mancala_grpc.pb.go src/protos/generated/mancala_pb.js src/protos/generated/mancala_grpc_pb.js
	cd src/protos
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mancala.proto
	protoc --js_out=import_style=commonjs,binary:generated --grpc_out=grpc_js:generated mancala.proto

backend: src/backend/server.go mancalaserver
	cd src/backend
	go build