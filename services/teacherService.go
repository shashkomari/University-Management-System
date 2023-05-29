package services

import (
	"fmt"
	"projects/course-work/models"
)

type teacherRepository interface {
	CreateTeacher(teacher models.TeacherData) (int64, error)
	GetTeacherById(id int64) (models.Teacher, error)
	UpdateTeacher(teacher models.Teacher) error
	DeleteTeacher(id int64) error
	GetTeachers() ([]models.Teacher, error)
}

type TeacherService struct {
	repository teacherRepository
}

func NewTeacherService(repository teacherRepository) *TeacherService {
	return &TeacherService{
		repository: repository,
	}
}

func (s *TeacherService) CreateTeacher(teacher models.TeacherData) (int64, error) {
	if err := checkName(&teacher.FirstName); err != nil {
		return -1, err
	}
	if err := checkName(&teacher.LastName); err != nil {
		return -1, err
	}
	if err := checkName(&teacher.MiddleName); teacher.MiddleName != "" && err != nil {
		return -1, err
	}

	id, err := s.repository.CreateTeacher(teacher)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *TeacherService) GetTeacherById(id int64) (models.Teacher, error) {
	var teacher models.Teacher
	var err error

	if id < 1 {
		return teacher, fmt.Errorf("bad request, check id")
	}

	teacher, err = s.repository.GetTeacherById(id)

	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func (s *TeacherService) DeleteTeacher(id int64) error {
	if id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.DeleteTeacher(id); err != nil {
		return err
	}

	return nil
}

func (s *TeacherService) UpdateTeacher(teacher models.Teacher) error {
	if err := checkName(&teacher.FirstName); err != nil {
		return err
	}
	if err := checkName(&teacher.LastName); err != nil {
		return err
	}
	if err := checkName(&teacher.MiddleName); teacher.MiddleName != "" && err != nil {
		return err
	}

	if teacher.Id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.UpdateTeacher(teacher); err != nil {
		return err
	}

	return nil
}

func (s *TeacherService) GetTeachers() ([]models.Teacher, error) {
	teachers, err := s.repository.GetTeachers()

	if err != nil {
		return teachers, err
	}

	return teachers, nil
}
