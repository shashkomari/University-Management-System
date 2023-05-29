package handlers

import (
	"log"
	"net/http"
	"projects/course-work/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lectureService interface {
	CreateLecture(lecture models.LectureData) (int64, error)
	GetLectureById(id int64) (models.Lecture, error)
	UpdateLecture(lecture models.Lecture) error
	DeleteLecture(id int64) error
	GetLectures() ([]models.Lecture, error)
}

type LectureHTTP struct {
	lectureService lectureService
}

func NewLectureHttp(lectureService lectureService) *LectureHTTP {
	return &LectureHTTP{lectureService: lectureService}
}

func (h *LectureHTTP) CreateLecture(c *gin.Context) {
	var newLecture models.LectureData

	if err := c.BindJSON(&newLecture); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := h.lectureService.CreateLecture(newLecture)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"NewItemID": id})
}

func (h *LectureHTTP) DeleteLecture(c *gin.Context) {
	lectureId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.lectureService.DeleteLecture(lectureId); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusGone, gin.H{})
}

func (h *LectureHTTP) GetLectureById(c *gin.Context) {
	lectureId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	lecture, err := h.lectureService.GetLectureById(lectureId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, lecture)
}

func (h *LectureHTTP) GetLectures(c *gin.Context) {
	lectures, err := h.lectureService.GetLectures()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusFound, lectures)
}

func (h *LectureHTTP) UpdateLecture(c *gin.Context) {
	var lecture models.Lecture

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	lecture.Id = id

	if err := c.BindJSON(&lecture); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := h.lectureService.UpdateLecture(lecture); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
