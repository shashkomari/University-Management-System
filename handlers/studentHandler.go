package handlers

import (
	"log"
	"net/http"
	"projects/course-work/models"

	"github.com/gin-gonic/gin"
)

type studentService interface {
	CreateStudent(student models.StudentData) (int, error)
	GetStudentById(id int) (models.Student, error)
	UpdateStudent(student models.Student) error
	DeleteStudent(id int) error
	GetStudents() ([]models.Student, error)
}

type StudentHTTP struct {
	studentService studentService
}

func NewStudentHttp(studentService studentService) *StudentHTTP {
	return &StudentHTTP{studentService: studentService}
}

func (h *StudentHTTP) CreateStudent(c *gin.Context) {
	var newStudent models.StudentData

	if err := c.BindJSON(&newStudent); err != nil {
		log.Println(err)
		return
	}

	id, err := h.studentService.CreateStudent(newStudent)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"NewItemID": id})
}

func (h *StudentHTTP) DeleteStudent(c *gin.Context) {
	studentId, err := getID(c)
	if err != nil {
		log.Println(err)
		return
	}

	if err := h.studentService.DeleteStudent(studentId); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusGone, gin.H{})
}

func (h *StudentHTTP) GetStudentById(c *gin.Context) {
	studentId, err := getID(c)
	if err != nil {
		log.Println(err)
		return
	}

	student, err := h.studentService.GetStudentById(studentId)
	if err != nil {
		log.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func (h *StudentHTTP) GetStudents(c *gin.Context) {
	students, err := h.studentService.GetStudents()
	if err != nil {
		log.Println(err)
		return
	}

	c.IndentedJSON(http.StatusFound, students)
}

func (h *StudentHTTP) UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		log.Println(err)
		return
	}

	if err := h.studentService.UpdateStudent(student); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
