package dto

type CourseDto struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IDcategory  int    `json:"id_category"`
}

type CoursesDto []CourseDto
