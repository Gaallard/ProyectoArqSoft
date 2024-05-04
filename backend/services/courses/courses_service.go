package courses

import (
	"backend/domain/courses"
)

func CreateCourse(request courses.CourseRequest) courses.CourseResponse {
	//validar con la base de datos

	return courses.CourseResponse{
		Message: "Curso creado",
	}
}
