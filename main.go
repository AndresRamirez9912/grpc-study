package main

import (
	db "gRPC/src/DB"
	examProtos "gRPC/src/protos/exam"
	studentProtos "gRPC/src/protos/student"

	"gRPC/src/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create connection to the DB
	repo, err := db.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB", err)
	}

	// Create server (handler with the Proto methods)
	server := server.NewServer(repo, repo)

	s := grpc.NewServer()
	studentProtos.RegisterStudentServiceServer(s, server)
	examProtos.RegisterExamServiceServer(s, server)

	// Add Reflection connection
	reflection.Register(s)

	// Create connection
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal("Error creating the listener", err)
	}

	log.Println("Server starting in port :5060")
	err = s.Serve(list)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
