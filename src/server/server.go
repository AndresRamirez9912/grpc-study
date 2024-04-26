package server

import (
	"context"
	"gRPC/src/models"
	examProtos "gRPC/src/protos/exam"
	studentProtos "gRPC/src/protos/student"
	"gRPC/src/repository"
	"io"
)

type Server struct {
	repo repository.Repository
	studentProtos.UnimplementedStudentServiceServer
	examProtos.UnimplementedExamServiceServer
}

func NewServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) GetStudent(ctx context.Context, request *studentProtos.StudentRequest) (*studentProtos.Student, error) {
	// Make the DB operation
	student, err := s.repo.GetStudent(ctx, request.GetId())
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
	err := s.repo.CreateStudent(ctx, newStudent)
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

	err := s.repo.CreateExam(ctx, newExam)
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
	exam, err := s.repo.GetExam(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	// Send response
	return &examProtos.Exam{
		Id:   exam.Id,
		Name: exam.Name,
	}, nil
}

func (s *Server) CreateQuestion(stream examProtos.ExamService_CreateQuestionServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF { // Check if the stream finish
			return stream.SendAndClose(&examProtos.CreateQuestionResponse{
				Ok: true,
			})
		}

		question := &models.Question{
			Id:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			ExamId:   msg.GetExamId(),
		}

		err = s.repo.CreateQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&examProtos.CreateQuestionResponse{
				Ok: false,
			})
		}
	}
}
