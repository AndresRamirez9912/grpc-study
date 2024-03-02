package db

import (
	"context"
	"database/sql"
	"gRPC/src/models"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// Ping the Connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the DB ", err)
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (repo *PostgresRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id,name,age) VALUES ($1,$2,$3)", student.Id, student.Name, student.Age)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	row, err := repo.db.QueryContext(ctx, "SELECT * FROM students WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal("Error closing the query", err)
		}
	}()

	student := &models.Student{}
	for row.Next() {
		err := row.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
	}
	return student, nil
}

func (repo *PostgresRepository) CreateExam(ctx context.Context, test *models.Exam) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO exams (id,name) VALUES ($1,$2)", test.Id, test.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) GetExam(ctx context.Context, id string) (*models.Exam, error) {
	row, err := repo.db.QueryContext(ctx, "SELECT * FROM exams WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal("Error closing the query", err)
		}
	}()

	test := &models.Exam{}
	for row.Next() {
		err := row.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
	}
	return test, nil
}
