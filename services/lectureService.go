package services

import (
	"fmt"
	"projects/course-work/models"
	"strconv"
	"strings"
)

type lectureRepository interface {
	CreateLecture(lecture models.LectureData) (int64, error)
	GetLectureById(id int64) (models.Lecture, error)
	UpdateLecture(lecture models.Lecture) error
	DeleteLecture(id int64) error
	GetLectures() ([]models.Lecture, error)
}

type LectureService struct {
	repository lectureRepository
}

func NewLectureService(repository lectureRepository) *LectureService {
	return &LectureService{
		repository: repository,
	}
}

func (s *LectureService) CreateLecture(lecture models.LectureData) (int64, error) {
	if err := checkDate(lecture.Date); err != nil {
		return -1, err
	}

	if lecture.SubjectId < 1 || lecture.TeacherId < 1 {
		return -1, fmt.Errorf("bad request, check id of subject or teacher")
	}

	id, err := s.repository.CreateLecture(lecture)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *LectureService) GetLectureById(id int64) (models.Lecture, error) {
	var lecture models.Lecture
	var err error

	if id < 1 {
		return lecture, fmt.Errorf("bad request, check id")
	}

	lecture, err = s.repository.GetLectureById(id)
	if err != nil {
		return lecture, err
	}

	return lecture, nil
}

func (s *LectureService) DeleteLecture(id int64) error {
	if id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.DeleteLecture(id); err != nil {
		return err
	}

	return nil
}

func (s *LectureService) UpdateLecture(lecture models.Lecture) error {
	if err := checkDate(lecture.Date); err != nil {
		return err
	}

	if lecture.Id < 1 {
		return fmt.Errorf("bad request, check id")
	}

	if err := s.repository.UpdateLecture(lecture); err != nil {
		return err
	}

	return nil
}

func (s *LectureService) GetLectures() ([]models.Lecture, error) {
	lectures, err := s.repository.GetLectures()

	if err != nil {
		return lectures, err
	}

	return lectures, nil
}

func checkDate(lectureDate string) error {
	stringNums := strings.Split(lectureDate, "-")

	if len(stringNums) != 3 {
		return fmt.Errorf("bad request, check date")
	}

	var intNums = []int{}

	for _, i := range stringNums {
		j, err := strconv.Atoi(i)
		if err != nil {
			return err
		}
		intNums = append(intNums, j)
	}
	fmt.Println(intNums)

	if intNums[0] < 1000 || intNums[0] > 9999 {
		return fmt.Errorf("bad request, check year of date")
	}

	if intNums[1] < 1 || intNums[1] > 12 {
		return fmt.Errorf("bad request, check month of date")
	}

	if intNums[2] < 1 || intNums[2] > 31 {
		return fmt.Errorf("bad request, check day of date")
	}

	return nil
}
