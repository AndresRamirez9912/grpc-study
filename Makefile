protoStudent: 
	protoc --proto_path=src/protos/student --go_out=. --go-grpc_out=. src/protos/student/student.proto

protoExam: 
	protoc --proto_path=src/protos/exam --go_out=. --go-grpc_out=. exam.proto

postgres:
	docker compose up
