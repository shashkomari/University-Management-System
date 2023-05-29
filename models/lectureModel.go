package models

type Lecture struct {
	Id        int64  `json:"id"`
	Date      string `json:"date"`
	SubjectId int64  `json:"subject_id"`
	TeacherId int64  `json:"teacher_id"`
}

type LectureData struct {
	Date      string `json:"date"`
	SubjectId int64  `json:"subject_id"`
	TeacherId int64  `json:"teacher_id"`
}
