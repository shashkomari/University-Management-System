package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"projects/course-work/handlers"
	"projects/course-work/repositories"
	"projects/course-work/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.Static("/css", "./css/")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/CourseWork")
	if err != nil {
		panic(err.Error())
	}

	lectureRepository := repositories.NewLectureRepository(db)
	lectureServices := services.NewLectureService(lectureRepository)
	lectureHandlers := handlers.NewLectureHttp(lectureServices)

	studentRepository := repositories.NewStudentRepository(db)
	studentServices := services.NewStudentService(studentRepository)
	studentHandlers := handlers.NewStudentHttp(studentServices)

	subjectRepository := repositories.NewSubjectRepository(db)
	subjectServices := services.NewSubjectService(subjectRepository)
	subjectHandlers := handlers.NewSubjectHttp(subjectServices)

	teacherRepository := repositories.NewTeacherRepository(db)
	teacherServices := services.NewTeacherService(teacherRepository)
	teacherHandlers := handlers.NewTeacherHttp(teacherServices)

	r.GET("/api/lectures", lectureHandlers.GetLectures)
	r.GET("/api/lecture/:id", lectureHandlers.GetLectureById)
	r.POST("/api/lecture", lectureHandlers.CreateLecture)
	r.PUT("/api/lecture/:id", lectureHandlers.UpdateLecture)
	r.DELETE("/api/lecture/:id", lectureHandlers.DeleteLecture)

	r.GET("/api/students", studentHandlers.GetStudents)
	r.GET("/api/student/:id", studentHandlers.GetStudentById)
	r.POST("/api/student", studentHandlers.CreateStudent)
	r.PUT("/api/student/:id", studentHandlers.UpdateStudent)
	r.DELETE("/api/student/:id", studentHandlers.DeleteStudent)

	r.GET("/api/subjects", subjectHandlers.GetSubjects)
	r.GET("/api/subject/:id", subjectHandlers.GetSubjectById)
	r.POST("/api/subject", subjectHandlers.CreateSubject)
	r.PUT("/api/subject/:id", subjectHandlers.UpdateSubject)
	r.DELETE("/api/subject/:id", subjectHandlers.DeleteSubject)

	r.GET("/api/teachers", teacherHandlers.GetTeachers)
	r.GET("/api/teacher/:id", teacherHandlers.GetTeacherById)
	r.POST("/api/teacher", teacherHandlers.CreateTeacher)
	r.PUT("/api/teacher/:id", teacherHandlers.UpdateTeacher)
	r.DELETE("/api/teacher/:id", teacherHandlers.DeleteTeacher)

	r.GET("/lectures", func(c *gin.Context) {
		c.HTML(http.StatusOK, "lectures.html", gin.H{})
	})
	r.GET("/students", func(c *gin.Context) {
		c.HTML(http.StatusOK, "students.html", gin.H{})
	})
	r.GET("/subjects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "subjects.html", gin.H{})
	})
	r.GET("/teachers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "teachers.html", gin.H{})
	})
	r.Run()
}
