package repositories

import (
	"database/sql"
	"fmt"
	"projects/course-work/models"
)

func NewSubjectRepository(db *sql.DB) *SubjectRepository {
	return &SubjectRepository{
		db: db,
	}
}

type SubjectRepository struct {
	db *sql.DB
}

func (r *SubjectRepository) GetSubjectById(id int64) (models.Subject, error) {
	subject := models.Subject{}

	err := r.db.QueryRow("SELECT * FROM subjects WHERE id = ?", id).Scan(&subject.Id, &subject.Name, &subject.Course, &subject.Semester)
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (r *SubjectRepository) UpdateSubject(subject models.Subject) error {
	res, err := r.db.Exec("UPDATE subjects SET name = ?, course = ?, semester = ? WHERE id = ?", subject.Name, subject.Course, subject.Semester, subject.Id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *SubjectRepository) DeleteSubject(id int64) error {
	res, err := r.db.Exec("DELETE FROM subjects WHERE id = ?", id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *SubjectRepository) CreateSubject(subject models.SubjectData) (int64, error) {
	var id int64

	res, err := r.db.Exec("INSERT INTO subjects (name, course, semester) VALUES (?, ?, ?)", subject.Name, subject.Course, subject.Semester)
	if err != nil {
		return id, err
	}

	return res.LastInsertId()
}

func (r *SubjectRepository) GetSubjects() ([]models.Subject, error) {
	var subjects []models.Subject

	rows, err := r.db.Query("SELECT * FROM subjects")

	if err != nil {
		return []models.Subject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		subject := models.Subject{}
		if err = rows.Scan(&subject.Id, &subject.Name, &subject.Course, &subject.Semester); err != nil && err != sql.ErrNoRows {
			return []models.Subject{}, err
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}
