package main

import (
	db "gRPC/src/DB"
	"gRPC/src/protos"
	"gRPC/src/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal("Error creating the listener", err)
	}

	// Create connection to the DB
	repo, err := db.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB", err)
	}

	server := server.NewStudentServer(repo)

	s := grpc.NewServer()
	protos.RegisterStudentServiceServer(s, server)

	// Add Reflection connection
	reflection.Register(s)

	log.Println("Server starting in port :5060")
	err = s.Serve(list)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
