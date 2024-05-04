package courses

type CourseRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CourseResponse struct {
	Message string `json:"message"`
}
