package repositories

import (
	"fmt"
	"projects/course-work/models"
)

func NewLocalLectureRepository() *LocalLectureRepository {
	db := []models.Lecture{
		{Id: 1, Date: "2021 01 20", SubjectId: 1, TeacherId: 56},
		{Id: 2, Date: "2020 03 1", SubjectId: 2, TeacherId: 17},
		{Id: 3, Date: "2023 07 12", SubjectId: 3, TeacherId: 39},
	}
	return &LocalLectureRepository{
		db:         db,
		id_counter: len(db) + 1,
	}
}

type LocalLectureRepository struct {
	db         []models.Lecture
	id_counter int
}

func (r *LocalLectureRepository) GetLectureById(id int) (models.Lecture, error) {
	lecture := models.Lecture{}

	for i := 0; i < len(r.db); i++ {
		if r.db[i].Id == id {
			return r.db[i], nil
		}
	}

	return lecture, fmt.Errorf("not found data")
}

func (r *LocalLectureRepository) UpdateLecture(lecture models.Lecture) error {

	for i := 0; i < len(r.db); i++ {
		if r.db[i].Id == lecture.Id {
			r.db[i].Date = lecture.Date
			r.db[i].SubjectId = lecture.SubjectId
			r.db[i].TeacherId = lecture.TeacherId
			return nil
		}
	}

	return fmt.Errorf("data not found")
}

func (r *LocalLectureRepository) DeleteLecture(id int) error {
	for i := 0; i < len(r.db); i++ {
		if r.db[i].Id == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("not found data")
}

func (r *LocalLectureRepository) CreateLecture(lectureData models.LectureData) (int, error) {
	var lecture models.Lecture
	lecture.Id = r.id_counter
	lecture.Date = lectureData.Date
	lecture.SubjectId = lectureData.SubjectId
	lecture.TeacherId = lectureData.TeacherId
	r.id_counter++

	r.db = append(r.db, lecture)

	id := lecture.Id

	return id, nil
}

func (r *LocalLectureRepository) GetLectures() ([]models.Lecture, error) {
	return r.db, nil
}
