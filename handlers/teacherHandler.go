package handlers

import (
	"projects/course-work/models"
)

type teacherService interface {
	CreateTeacher(teacher models.Teacher) (models.Teacher, error)
	DeleteTeacher(id models.TeacherId) (models.Teacher, error)
	GetTeacher(id models.TeacherId) (models.Teacher, error)
	UpdateTeacher(id models.TeacherId) (models.Teacher, error)
}

type TeacherHTTP struct {
	teacherService teacherService
}

func NewTeacherHttp(teacherService teacherService) *TeacherHTTP {
	return &TeacherHTTP{teacherService: teacherService}
}
