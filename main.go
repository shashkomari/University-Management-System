package main

import (
	"database/sql"

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

	r.GET("/lectures", lectureHandlers.GetLectures)
	r.GET("/lecture/:id", lectureHandlers.GetLectureById)
	r.POST("/lecture", lectureHandlers.CreateLecture)
	r.PUT("/lecture/:id", lectureHandlers.UpdateLecture)
	r.DELETE("/lecture/:id", lectureHandlers.DeleteLecture)

	r.GET("/students", studentHandlers.GetStudents)
	r.GET("/student/:id", studentHandlers.GetStudentById)
	r.POST("/student", studentHandlers.CreateStudent)
	r.PUT("/student/:id", studentHandlers.UpdateStudent)
	r.DELETE("/student/:id", studentHandlers.DeleteStudent)

	r.GET("/subjects", subjectHandlers.GetSubjects)
	r.GET("/subject/:id", subjectHandlers.GetSubjectById)
	r.POST("/subject", subjectHandlers.CreateSubject)
	r.PUT("/subject/:id", subjectHandlers.UpdateSubject)
	r.DELETE("/subject/:id", subjectHandlers.DeleteSubject)

	r.GET("/teachers", teacherHandlers.GetTeachers)
	r.GET("/teacher/:id", teacherHandlers.GetTeacherById)
	r.POST("/teacher", teacherHandlers.CreateTeacher)
	r.PUT("/teacher/:id", teacherHandlers.UpdateTeacher)
	r.DELETE("/teacher/:id", teacherHandlers.DeleteTeacher)

	r.Run()
}
