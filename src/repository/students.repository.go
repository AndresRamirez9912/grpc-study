package repository

import (
	"context"
	"gRPC/src/models"
)

type StudentRepository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) error
}

type ExamRepository interface {
	GetExam(ctx context.Context, id string) (*models.Exam, error)
	CreateExam(ctx context.Context, exam *models.Exam) error
}

type QuestionRepository interface {
	CreateQuestion(ctx context.Context, question *models.Question) error
}

type Repository interface {
	StudentRepository
	ExamRepository
	QuestionRepository
}
