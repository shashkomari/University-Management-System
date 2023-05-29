package repositories

import (
	"database/sql"
	"fmt"
	"projects/course-work/models"
)

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{
		db: db,
	}
}

type TeacherRepository struct {
	db *sql.DB
}

func (r *TeacherRepository) GetTeacherById(id int64) (models.Teacher, error) {
	teacher := models.Teacher{}

	err := r.db.QueryRow("SELECT * FROM teachers WHERE id = ?", id).Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.MiddleName)
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func (r *TeacherRepository) UpdateTeacher(teacher models.Teacher) error {
	res, err := r.db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, middle_name = ? WHERE id = ?", teacher.FirstName, teacher.LastName, teacher.MiddleName, teacher.Id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *TeacherRepository) DeleteTeacher(id int64) error {
	res, err := r.db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *TeacherRepository) CreateTeacher(teacher models.TeacherData) (int64, error) {
	var id int64

	res, err := r.db.Exec("INSERT INTO teachers (first_name, last_name, middle_name) VALUES (?, ?, ?)", teacher.FirstName, teacher.LastName, teacher.MiddleName)
	if err != nil {
		return id, err
	}

	return res.LastInsertId()
}

func (r *TeacherRepository) GetTeachers() ([]models.Teacher, error) {
	var teachers []models.Teacher

	rows, err := r.db.Query("SELECT * FROM teachers")

	if err != nil {
		return []models.Teacher{}, err
	}
	defer rows.Close()

	for rows.Next() {
		teacher := models.Teacher{}
		if err = rows.Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.MiddleName); err != nil && err != sql.ErrNoRows {
			return []models.Teacher{}, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}
