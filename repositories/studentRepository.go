package repositories

import (
	"database/sql"
	"fmt"
	"projects/course-work/models"
)

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

type StudentRepository struct {
	db *sql.DB
}

func (r *StudentRepository) GetStudentById(id int64) (models.Student, error) {
	student := models.Student{}

	err := r.db.QueryRow("SELECT * FROM students WHERE id = ?", id).Scan(&student.Id, &student.FirstName, &student.LastName, &student.MiddleName, &student.GroupNum)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (r *StudentRepository) UpdateStudent(student models.Student) error {
	res, err := r.db.Exec("UPDATE students SET first_name = ?, last_name = ?, middle_name = ?, group_num = ? WHERE id = ?", student.FirstName, student.LastName, student.MiddleName, student.GroupNum, student.Id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *StudentRepository) DeleteStudent(id int64) error {
	res, err := r.db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *StudentRepository) CreateStudent(student models.StudentData) (int64, error) {
	var id int64

	res, err := r.db.Exec("INSERT INTO students (first_name, last_name, middle_name, group_num) VALUES (?, ?, ?, ?)", student.FirstName, student.LastName, student.MiddleName, student.GroupNum)
	if err != nil {
		return id, err
	}

	return res.LastInsertId()
}

func (r *StudentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student

	rows, err := r.db.Query("SELECT * FROM students")

	if err != nil {
		return []models.Student{}, err
	}
	defer rows.Close()

	for rows.Next() {
		student := models.Student{}
		if err = rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.MiddleName, &student.GroupNum); err != nil && err != sql.ErrNoRows {
			return []models.Student{}, err
		}
		students = append(students, student)
	}

	return students, nil
}
