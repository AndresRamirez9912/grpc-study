fixPath:
	PATH="${PATH}:${HOME}/go/bin"

protoStudent: 
	protoc --proto_path=src/protos/student --go_out=. --go-grpc_out=. src/protos/student/student.proto

protoExam: 
	protoc --proto_path=src/protos/exam --go_out=. --go-grpc_out=. src/protos/exam/exam.proto

postgres:
	docker compose up
