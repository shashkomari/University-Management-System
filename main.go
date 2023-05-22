package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"projects/course-work/handlers"
	//"projects/course-work/models"
	"projects/course-work/repositories"
	"projects/course-work/services"

	_ "github.com/go-sql-driver/mysql"
)

// type User struct {
// 	Title string
// }

// var hello = []User{
// 	{"Hello!"},
// }

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

	r.GET("/lectures", lectureHandlers.GetLectures)
	r.GET("/lecture/:id", lectureHandlers.GetLectureById)
	r.POST("/lecture", lectureHandlers.CreateLecture)
	r.PUT("/lecture/:id", lectureHandlers.UpdateLecture)
	r.DELETE("/lecture/:id", lectureHandlers.DeleteLecture)
	// r.GET("/", func(c *gin.Context) {
	// 	//c.Data(200, "application/json; charset=utf-8", []byte("WELCOME!"))
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "Main website",
	// 	})
	// })

	//r.POST("/", lectureHandlers.CreateLecture)

	// r.GET("/get", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, hello)
	// })

	// r.POST("/post", func(c *gin.Context) {
	// 	user := c.DefaultPostForm("user", "unknown")
	// 	c.String(200, "hello %s", user)
	// })
	r.Run()
}
