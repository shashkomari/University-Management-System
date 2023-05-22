package repositories

import (
	"database/sql"
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

func (r *LectureRepository) GetLectureById(id int) (models.Lecture, error) {
	lecture := models.Lecture{}

	err := r.db.QueryRow("SELECT * FROM lectures WHERE id = ?", id).Scan(&lecture.Id, &lecture.Date, &lecture.SubjectId, &lecture.TeacherId)

	if err != nil && err != sql.ErrNoRows {
		return models.Lecture{}, err
	}
	return lecture, nil
}

func (r *LectureRepository) UpdateLecture(lecture models.Lecture) error {
	err := r.db.QueryRow("UPDATE lectures SET date = ?, subject_id = ?, teacher_id = ? WHERE id = ?", lecture.Date, lecture.SubjectId, lecture.TeacherId, lecture.Id).Err()

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *LectureRepository) DeleteLecture(id int) error {
	_, err := r.db.Exec("DELETE FROM lectures WHERE id = ?", id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *LectureRepository) CreateLecture(lecture models.LectureData) (int, error) {
	var id int

	_, err := r.db.Exec("INSERT INTO lectures (date, subject_id, teacher_id) VALUES (?, ?, ?)", lecture.Date, lecture.SubjectId, lecture.TeacherId)
	if err != nil {
		return id, err
	}

	err = r.db.QueryRow("SELECT last_insert_id()").Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
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
