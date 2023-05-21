package handlers

import (
	"projects/course-work/models"
)

type subjectService interface {
	CreateSubject(subject models.Subject) (models.Subject, error)
	DeleteSubject(id models.SubjectId) (models.Subject, error)
	GetSubject(id models.SubjectId) (models.Subject, error)
	UpdateSubject(id models.SubjectId) (models.Subject, error)
}

type SubjectHTTP struct {
	subjectService subjectService
}

func NewSubjectHttp(subjectService subjectService) *SubjectHTTP {
	return &SubjectHTTP{subjectService: subjectService}
}
