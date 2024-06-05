package courses

import (
	dtoCourses "backend/dto/courses"
	service "backend/services/courses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCourse(context *gin.Context) {
	var courseRequest dtoCourses.CourseDto
	context.BindJSON(&courseRequest)
	response, err := service.CreateCourse(courseRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, response)
}

func UpdateCourse(context *gin.Context) {
	var courseRequest dtoCourses.CourseDto
	context.BindJSON(&courseRequest)
	response, err := service.UpdateCourse(courseRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func DeleteCourse(context *gin.Context) {
	var courseRequest dtoCourses.CourseDto
	context.BindJSON(&courseRequest)
	err := service.DeleteCourse(courseRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, "Curso eliminado")
}

func GetCoursesByUser(context *gin.Context) {
	idUser, _ := strconv.Atoi(context.Param("idUser"))
	response, err := service.GetCoursesByUser(idUser)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}
