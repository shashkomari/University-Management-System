package services

import (
	"fmt"
	"projects/course-work/models"
)

type studentRepository interface {
	CreateStudent(student models.StudentData) (int64, error)
	GetStudentById(id int64) (models.Student, error)
	UpdateStudent(student models.Student) error
	DeleteStudent(id int64) error
	GetStudents() ([]models.Student, error)
}

type StudentService struct {
	repository studentRepository
}

func NewStudentService(repository studentRepository) *StudentService {
	return &StudentService{
		repository: repository,
	}
}

func (s *StudentService) CreateStudent(student models.StudentData) (int64, error) {
	if err := checkName(&student.FirstName); err != nil {
		return -1, err
	}
	if err := checkName(&student.LastName); err != nil {
		return -1, err
	}
	if err := checkName(&student.MiddleName); student.MiddleName != "" && err != nil {
		return -1, err
	}

	if student.GroupNum < 100 || student.GroupNum > 999 {
		return -1, fmt.Errorf("bad request, check group number (it should be a three-digit number)")
	}

	id, err := s.repository.CreateStudent(student)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *StudentService) GetStudentById(id int64) (models.Student, error) {
	var student models.Student
	var err error

	if id < 1 {
		return student, fmt.Errorf("bad request, check id")
	}

	student, err = s.repository.GetStudentById(id)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *StudentService) DeleteStudent(id int64) error {
	if id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.DeleteStudent(id); err != nil {
		return err
	}

	return nil
}

func (s *StudentService) UpdateStudent(student models.Student) error {
	if err := checkName(&student.FirstName); err != nil {
		return err
	}
	if err := checkName(&student.LastName); err != nil {
		return err
	}
	if err := checkName(&student.MiddleName); student.MiddleName != "" && err != nil {
		return err
	}

	if student.GroupNum < 100 || student.GroupNum > 999 {
		return fmt.Errorf("bad request, check group number (it should be a three-digit number)")
	}

	if student.Id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.UpdateStudent(student); err != nil {
		return err
	}

	return nil
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	students, err := s.repository.GetStudents()

	if err != nil {
		return students, err
	}

	return students, nil
}
