package server

import (
	"context"
	"gRPC/src/models"
	"gRPC/src/protos"
	"gRPC/src/repository"
)

type Server struct {
	repo repository.Repository
	protos.UnimplementedStudentServiceServer 
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, request *protos.StudentRequest) (*protos.Student, error) {
	student, err := s.repo.GetStudent(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &protos.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) CreateStudent(ctx context.Context, student *protos.Student) (*protos.StudentResponse, error) {
	newStudent := &models.Student{
		Id:   student.GetId(),
		Name: student.GetName(),
		Age:  student.GetAge(),
	}
	err := s.repo.CreateStudent(ctx, newStudent)
	if err != nil {
		return nil, err
	}
	return &protos.StudentResponse{
		Id: newStudent.Id,
	}, nil
}
