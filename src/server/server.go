package server

import (
	"context"
	"gRPC/src/models"
	examProtos "gRPC/src/protos/exam"
	studentProtos "gRPC/src/protos/student"
	"gRPC/src/repository"
)

type Server struct {
	studentRepo repository.StudentRepository
	examRepo    repository.ExamRepository
	studentProtos.UnimplementedStudentServiceServer
	examProtos.UnimplementedExamServiceServer
}

func NewServer(studentRepo repository.StudentRepository, examRepo repository.ExamRepository) *Server {
	return &Server{studentRepo: studentRepo, examRepo: examRepo}
}

func (s *Server) GetStudent(ctx context.Context, request *studentProtos.StudentRequest) (*studentProtos.Student, error) {
	// Make the DB operation
	student, err := s.studentRepo.GetStudent(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	// Send data
	return &studentProtos.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) CreateStudent(ctx context.Context, student *studentProtos.Student) (*studentProtos.StudentResponse, error) {
	// Make the DB operation
	newStudent := &models.Student{
		Id:   student.GetId(),
		Name: student.GetName(),
		Age:  student.GetAge(),
	}
	err := s.studentRepo.CreateStudent(ctx, newStudent)
	if err != nil {
		return nil, err
	}

	// Send data
	return &studentProtos.StudentResponse{
		Id: newStudent.Id,
	}, nil
}

func (s *Server) CreateExam(ctx context.Context, exam *examProtos.CreateExamRequest) (*examProtos.CreateExamResponse, error) {
	// Make the DB operation
	newExam := &models.Exam{
		Id:   exam.GetId(),
		Name: exam.GetName(),
	}

	err := s.examRepo.CreateExam(ctx, newExam)
	if err != nil {
		return nil, err
	}

	// Send response
	return &examProtos.CreateExamResponse{
		Id: newExam.Id,
	}, nil
}

func (s *Server) GetExam(ctx context.Context, request *examProtos.GetExamRequest) (*examProtos.Exam, error) {
	// Make the DB operation
	exam, err := s.examRepo.GetExam(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	// Send response
	return &examProtos.Exam{
		Id:   exam.Id,
		Name: exam.Name,
	}, nil
}
