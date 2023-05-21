package repositories

import (
	"database/sql"
)

func NewLectureRepository(db *sql.DB) *LectureRepository {
	return &LectureRepository{
		db: db,
	}
}

type LectureRepository struct {
	db *sql.DB
}

/*
func (r *LectureRepository) GetLectureById(id models.LectureId) (models.Lecture, error) {
	lecture := models.Lecture{}

	err := r.db.QueryRow("SELECT * FROM lectures WHERE id=$1", id).Scan(&lecture.Id, &lecture.Date, &lecture.SubjectId, &lecture.TeacherId)

	if err != nil && err != sql.ErrNoRows {
		return models.Lecture{}, err
	}
	return lecture, nil
}

func (r *LectureRepository) UpdateLecture(id models.LectureId, Lecture models.Lecture) error {
	err := r.db.QueryRow("UPDATE lectures SET id = $1, date = $2, subject_id = $3, teacher_id = $4 WHERE id = $5", Lecture.Id, Lecture.Date, Lecture.SubjectId, Lecture.TeacherId, id).Err()

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *LectureRepository) DeleteLecture(id models.LectureId) error {
	_, err := r.db.Exec("DELETE FROM lectures WHERE id = $1", id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *LectureRepository) CreateLecture(lecture models.LectureData) (models.LectureId, error) {
	var id models.LectureId
	err := r.db.QueryRow("INSERT INTO lectures (date, subject_id, teacher_id) VALUES ($1, $2, $3) RETURNING id", lecture.Date, lecture.SubjectId, lecture.TeacherId).Scan(&id.Id)

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
*/

///////// ONLY FOR TEST /////////

/////////////////////////////////
