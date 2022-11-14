package repository

import (
	"context"

	"github.com/adrisongomez/grpc_golang/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error

	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error

	GetQuestion(ctx context.Context, id string) (*models.Question, error)
	SetQuestion(ctx context.Context, question *models.Question) error
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)

	GetStudentPerTest(ctx context.Context, id string) ([]*models.Student, error)
	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	return implementation.GetQuestion(ctx, id)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enrollment)
}

func GetStudentPerTest(ctx context.Context, id string) ([]*models.Student, error) {
	return implementation.GetStudentPerTest(ctx, id)
}

func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionsPerTest(ctx, testId)
}
