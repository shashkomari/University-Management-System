package handlers

import (
	"log"
	"net/http"
	"projects/course-work/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type studentService interface {
	CreateStudent(student models.StudentData) (int64, error)
	GetStudentById(id int64) (models.Student, error)
	UpdateStudent(student models.Student) error
	DeleteStudent(id int64) error
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
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := h.studentService.CreateStudent(newStudent)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"NewItemID": id})
}

func (h *StudentHTTP) DeleteStudent(c *gin.Context) {
	studentId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.studentService.DeleteStudent(studentId); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusGone, gin.H{})
}

func (h *StudentHTTP) GetStudentById(c *gin.Context) {
	studentId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	student, err := h.studentService.GetStudentById(studentId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func (h *StudentHTTP) GetStudents(c *gin.Context) {
	students, err := h.studentService.GetStudents()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusFound, students)
}

func (h *StudentHTTP) UpdateStudent(c *gin.Context) {
	var student models.Student

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	student.Id = id

	if err := c.BindJSON(&student); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.studentService.UpdateStudent(student); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
