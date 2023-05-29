package services

import (
	"fmt"
	"projects/course-work/models"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type subjectRepository interface {
	CreateSubject(subject models.SubjectData) (int64, error)
	GetSubjectById(id int64) (models.Subject, error)
	UpdateSubject(subject models.Subject) error
	DeleteSubject(id int64) error
	GetSubjects() ([]models.Subject, error)
}

type SubjectService struct {
	repository subjectRepository
}

func NewSubjectService(repository subjectRepository) *SubjectService {
	return &SubjectService{
		repository: repository,
	}
}

func (s *SubjectService) CreateSubject(subject models.SubjectData) (int64, error) {
	if err := checkSubjectName(&subject.Name); err != nil {
		return -1, err
	}

	if subject.Course < 1 || subject.Course > 5 {
		return -1, fmt.Errorf("bad request, check course: it cannot be zero or greater than 5")
	}

	if subject.Semester != 1 && subject.Semester != 2 {
		return -1, fmt.Errorf("bad request, check semester: it can only be 1 or 2")
	}

	id, err := s.repository.CreateSubject(subject)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *SubjectService) GetSubjectById(id int64) (models.Subject, error) {
	var subject models.Subject
	var err error

	if id < 1 {
		return subject, fmt.Errorf("bad request, check id")
	}

	subject, err = s.repository.GetSubjectById(id)

	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (s *SubjectService) DeleteSubject(id int64) error {
	if id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.DeleteSubject(id); err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) UpdateSubject(subject models.Subject) error {
	if err := checkSubjectName(&subject.Name); err != nil {
		return err
	}

	if subject.Course < 1 || subject.Course > 5 {
		return fmt.Errorf("bad request, check course: it cannot be zero or greater than 5")
	}

	if subject.Semester != 1 && subject.Semester != 2 {
		return fmt.Errorf("bad request, check semester: it can only be 1 or 2")
	}

	if subject.Id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.UpdateSubject(subject); err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) GetSubjects() ([]models.Subject, error) {
	subjects, err := s.repository.GetSubjects()

	if err != nil {
		return subjects, err
	}

	return subjects, nil
}

func checkSubjectName(name *string) error {

	for _, r := range *name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != ' ' {
			return fmt.Errorf("bad request, check subject name: it cannot contain special characters and nums")
		}
	}

	stringName := strings.Fields(*name)

	if len(stringName) > 3 {
		return fmt.Errorf("bad request, check subject name")
	}
	*name = ""
	for i := range stringName {
		stringName[i] = cases.Title(language.Und).String(stringName[i])
		stringName[i] = strings.ReplaceAll(stringName[i], " ", "")
		if stringName[i] != "" && i != len(stringName)-1 {
			*name = *name + stringName[i] + " "
		}
		if i == len(stringName)-1 {
			*name = *name + stringName[i]
		}
	}

	return nil
}
