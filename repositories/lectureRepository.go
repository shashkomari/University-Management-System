package repositories

import (
	"database/sql"
	"fmt"
	"projects/course-work/models"
)

func NewLectureRepository(db *sql.DB) *LectureRepository {
	return &LectureRepository{
		db: db,
	}
}

type LectureRepository struct {
	db *sql.DB
}

func (r *LectureRepository) GetLectureById(id int64) (models.Lecture, error) {
	lecture := models.Lecture{}

	err := r.db.QueryRow("SELECT * FROM lectures WHERE id = ?", id).Scan(&lecture.Id, &lecture.Date, &lecture.SubjectId, &lecture.TeacherId)
	if err != nil {
		return lecture, err
	}

	return lecture, nil
}

func (r *LectureRepository) UpdateLecture(lecture models.Lecture) error {
	res, err := r.db.Exec("UPDATE lectures SET date = ?, subject_id = ?, teacher_id = ? WHERE id = ?", lecture.Date, lecture.SubjectId, lecture.TeacherId, lecture.Id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *LectureRepository) DeleteLecture(id int64) error {
	res, err := r.db.Exec("DELETE FROM lectures WHERE id = ?", id)
	if err != nil {
		return err
	}
	if rf, _ := res.RowsAffected(); rf == 0 {
		return fmt.Errorf("id not found")
	}

	return nil
}

func (r *LectureRepository) CreateLecture(lecture models.LectureData) (int64, error) {
	var id int64

	res, err := r.db.Exec("INSERT INTO lectures (date, subject_id, teacher_id) VALUES (?, ?, ?)", lecture.Date, lecture.SubjectId, lecture.TeacherId)
	if err != nil {
		return id, err
	}

	return res.LastInsertId()
}

func (r *LectureRepository) GetLectures() ([]models.Lecture, error) {
	var lectures []models.Lecture

	rows, err := r.db.Query("SELECT * FROM lectures")

	if err != nil {
		return []models.Lecture{}, err
	}
	defer rows.Close()

	for rows.Next() {
		lecture := models.Lecture{}
		if err = rows.Scan(&lecture.Id, &lecture.Date, &lecture.SubjectId, &lecture.TeacherId); err != nil && err != sql.ErrNoRows {
			return []models.Lecture{}, err
		}
		lectures = append(lectures, lecture)
	}

	return lectures, nil
}
