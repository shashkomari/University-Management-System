package handlers

import (
	"log"
	"net/http"
	"projects/course-work/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type teacherService interface {
	CreateTeacher(teacher models.TeacherData) (int64, error)
	GetTeacherById(id int64) (models.Teacher, error)
	UpdateTeacher(teacher models.Teacher) error
	DeleteTeacher(id int64) error
	GetTeachers() ([]models.Teacher, error)
}

type TeacherHTTP struct {
	teacherService teacherService
}

func NewTeacherHttp(teacherService teacherService) *TeacherHTTP {
	return &TeacherHTTP{teacherService: teacherService}
}

func (h *TeacherHTTP) CreateTeacher(c *gin.Context) {
	var newTeacher models.TeacherData

	if err := c.BindJSON(&newTeacher); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := h.teacherService.CreateTeacher(newTeacher)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"NewItemID": id})
}

func (h *TeacherHTTP) DeleteTeacher(c *gin.Context) {
	teacherId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.teacherService.DeleteTeacher(teacherId); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusGone, gin.H{})
}

func (h *TeacherHTTP) GetTeacherById(c *gin.Context) {
	teacherId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	teacher, err := h.teacherService.GetTeacherById(teacherId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, teacher)
}

func (h *TeacherHTTP) GetTeachers(c *gin.Context) {
	teachers, err := h.teacherService.GetTeachers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusFound, teachers)
}

func (h *TeacherHTTP) UpdateTeacher(c *gin.Context) {
	var teacher models.Teacher

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	teacher.Id = id

	if err := c.BindJSON(&teacher); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.teacherService.UpdateTeacher(teacher); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
