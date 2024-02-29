package repository

import (
	"context"
	"gRPC/src/models"
)

var implementation Repository

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	CreateStudent(ctx context.Context, student *models.Student) error
}

func AssignStudentImplementation(repository Repository) {
	implementation = repository
}

func CreateStudent(ctx context.Context, student *models.Student) error {
	return implementation.CreateStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}
