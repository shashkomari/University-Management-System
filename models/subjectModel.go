package models

type Subject struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Course   int    `json:"course"`
	Semester int    `json:"semester"`
}

type SubjectData struct {
	Name     string `json:"name"`
	Course   int    `json:"course"`
	Semester int    `json:"semester"`
}
