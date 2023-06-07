package handlers

import (
	"log"
	"net/http"
	"projects/course-work/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type subjectService interface {
	CreateSubject(subject models.SubjectData) (int64, error)
	GetSubjectById(id int64) (models.Subject, error)
	UpdateSubject(subject models.Subject) error
	DeleteSubject(id int64) error
	GetSubjects() ([]models.Subject, error)
}

type SubjectHTTP struct {
	subjectService subjectService
}

func NewSubjectHttp(subjectService subjectService) *SubjectHTTP {
	return &SubjectHTTP{subjectService: subjectService}
}

func (h *SubjectHTTP) CreateSubject(c *gin.Context) {
	var newSubject models.SubjectData

	if err := c.BindJSON(&newSubject); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := h.subjectService.CreateSubject(newSubject)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"NewItemID": id})
}

func (h *SubjectHTTP) DeleteSubject(c *gin.Context) {
	subjectId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.subjectService.DeleteSubject(subjectId); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *SubjectHTTP) GetSubjectById(c *gin.Context) {
	subjectId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	subject, err := h.subjectService.GetSubjectById(subjectId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusFound, subject)
}

func (h *SubjectHTTP) GetSubjects(c *gin.Context) {
	subjects, err := h.subjectService.GetSubjects()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusFound, subjects)
}

func (h *SubjectHTTP) UpdateSubject(c *gin.Context) {
	var subject models.Subject

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	subject.Id = id

	if err := c.BindJSON(&subject); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.subjectService.UpdateSubject(subject); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
