package courses

import (
	dto "backend/dto/courses"
	"backend/middleware"
	service "backend/services/courses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCourse(context *gin.Context) {
	var courseRequest dto.CourseDto
	context.BindJSON(&courseRequest)
	response, err := service.CreateCourse(courseRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, response)
}

func UpdateCourse(context *gin.Context) {
	var courseRequest dto.CourseDto
	context.BindJSON(&courseRequest)
	response, err := service.UpdateCourse(courseRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func DeleteCourse(c *gin.Context) {
	idCourse, err := strconv.Atoi(c.Param("idCourse"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid idCourse"})
		return
	}
	response := service.DeleteCourse(idCourse)
	if response != nil {
		c.JSON(response.Status(), response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetCoursesByUser(context *gin.Context) {
	idUser, _ := middleware.GetUserIdByToken(context)
	response, err := service.GetCoursesByUser(idUser)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetCourses(context *gin.Context) {
	response, err := service.GetCourses()
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}
