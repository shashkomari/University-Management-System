package models

type Lecture struct {
	Id        int    `json:"id"`
	Date      string `json:"date"`
	SubjectId int    `json:"subject_id"`
	TeacherId int    `json:"teacher_id"`
}

type LectureData struct {
	Date      string `json:"date"`
	SubjectId int    `json:"subject_id"`
	TeacherId int    `json:"teacher_id"`
}
