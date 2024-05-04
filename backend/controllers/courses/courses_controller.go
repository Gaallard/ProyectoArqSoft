package courses

import (
	CoursesDomain "backend/domain/courses"
	CoursesService "backend/services/courses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(context *gin.Context) {
	var courseRequest CoursesDomain.CourseRequest
	context.BindJSON(&courseRequest)
	response := CoursesService.CreateCourse(courseRequest)
	context.JSON(http.StatusOK, response)
}
