package database

import (
	"context"
	"database/sql"

	"github.com/adrisongomez/grpc_golang/models"
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
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)",
		student.Id,
		student.Name,
		student.Age,
	)
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, name, age  FROM students WHERE id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}

	defer handleCloseRows(rows)

	return mapFromRowsToStudent(rows)
}

func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO tests (id, name) VALUES ($1, $2)",
		test.Id,
		test.Name,
	)
	return err
}

func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, name FROM tests WHERE id = $1",
		id,
	)

	if err != nil {
		return nil, err
	}

	defer handleCloseRows(rows)

	return mapFromRowsToTest(rows)
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO questions (id, test_id, questions, answer) VALUES ($1, $2, $3, $4)",
		question.Id,
		question.TestId,
		question.Question,
		question.Answer,
	)

	return err
}

func (repo *PostgresRepository) GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, test_id, questions, answer FROM questions WHERE id = $1",
		id,
	)

	if err != nil {
		return nil, err
	}

	defer handleCloseRows(rows)

	question := &models.Question{}

	for rows.Next() {
		err := rows.Scan(question.Id, question.TestId, question.Question, question.Answer)

		if err != nil {
			return nil, err
		}
	}

	return question, nil
}

func (repo *PostgresRepository) SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO enrollments (test_id, student_id) VALUES ($1, $2)",
		enrollment.TestId,
		enrollment.StudentId,
	)

	return err
}

func (repo *PostgresRepository) GetStudentPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)",
		testId,
	)

	if err != nil {
		return nil, err
	}

	defer handleCloseRows(rows)

	students := []*models.Student{}

	for rows.Next() {
		student := models.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err == nil {
			students = append(students, &student)
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (repo *PostgresRepository) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, questions, answer, test_id FROM questions WHERE test_id = $1",
		testId,
	)

	if err != nil {
		return nil, err
	}

	defer handleCloseRows(rows)

	questions := []*models.Question{}

	for rows.Next() {
		question := models.Question{}
		if err = rows.Scan(&question.Id, &question.Question, &question.Answer, &question.TestId); err == nil {
			questions = append(questions, &question)
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

    return questions, nil
}
