package courses

import (
	client "backend/clients/courses"
	dto "backend/dto/courses"
	e "backend/errors"
	model "backend/models"
)

func CreateCourse(request dto.CourseDto) (dto.CourseDto, e.ApiError) {
	var course model.Course

	course.Name = request.Name
	course.Description = request.Description
	course.Id_category = request.IDcategory

	course, err := client.CreateCourse(course)
	if err != nil {
		return request, e.NewBadRequestApiError("Error al crear curso")
	}

	request.ID = course.Id_course

	return request, nil
}

func UpdateCourse(request dto.CourseDto) (dto.CourseDto, e.ApiError) {
	var course model.Course

	course.Id_course = request.ID
	course.Name = request.Name
	course.Description = request.Description
	course.Id_category = request.IDcategory

	course, err := client.UpdateCourse(course)
	if err != nil {
		return request, e.NewBadRequestApiError("Error al actualizar curso")
	}

	return request, nil
}

func DeleteCourse(request dto.CourseDto) e.ApiError {
	var course model.Course

	course.Id_course = request.ID
	err := client.DeleteCourse(course)
	if err != nil {
		return e.NewBadRequestApiError("Error al eliminar curso")
	}

	return nil
}

func GetCoursesByUser(idUser int) ([]dto.CourseDto, e.ApiError) {
	var courses []model.Course

	courses, err := client.GetCoursesByUser(idUser)
	if err != nil {
		return nil, e.NewBadRequestApiError("Error al buscar cursos")
	}

	var coursesDto []dto.CourseDto
	for _, course := range courses {
		coursesDto = append(coursesDto, dto.CourseDto{
			ID:          course.Id_course,
			Name:        course.Name,
			Description: course.Description,
			IDcategory:  course.Id_category,
		})
	}

	return coursesDto, nil
}

func GetCourses() ([]dto.CourseDto, e.ApiError) {
	var courses []model.Course

	courses, err := client.GetCourses()
	if err != nil {
		return nil, e.NewBadRequestApiError("Error al buscar cursos")
	}

	var coursesDto []dto.CourseDto
	for _, course := range courses {
		coursesDto = append(coursesDto, dto.CourseDto{
			ID:          course.Id_course,
			Name:        course.Name,
			Description: course.Description,
			IDcategory:  course.Id_category,
		})
	}

	return coursesDto, nil
}
