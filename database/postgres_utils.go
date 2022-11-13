package database

import (
	"database/sql"
	"log"

	"github.com/adrisongomez/grpc_golang/models"
)

func mapFromRowsToStudent(rows *sql.Rows) (*models.Student, error) {
	student := models.Student{}

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
	}
	return &student, nil
}

func handleCloseRows(rows *sql.Rows) {
	err := rows.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func mapFromRowsToTest(rows *sql.Rows) (*models.Test, error) {
	test := models.Test{}

	for rows.Next() {
		if err := rows.Scan(&test.Id, &test.Name); err != nil {
			return nil, err
		}
	}

	return &test, nil
}
