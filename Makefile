protoEstudent: 
	protoc --proto_path=src/protos --go_out=. --go-grpc_out=. src/protos/student.proto

postgres:
	docker compose up
