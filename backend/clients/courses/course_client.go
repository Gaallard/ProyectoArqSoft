package courses

import (
	"backend/models"

	e "backend/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func CreateCourse(course models.Course) (models.Course, e.ApiError) {
	result := Db.Create(&course)
	log.Info("Curso creado con exito")
	if result.Error != nil {
		log.Error("Error al crear el curso")
		log.Error(result.Error)
		return course, e.NewBadRequestApiError("Error al crear curso")
	}

	return course, nil
}

func UpdateCourse(course models.Course) (models.Course, e.ApiError) {
	result := Db.Model(&course).Updates(models.Course{Name: course.Name, Description: course.Description, Id_category: course.Id_category}).Where("Id_course = ?", course.Id_course)
	if result.Error != nil {
		log.Error("Error al actualizar el curso")
		log.Error(result.Error)
		return course, e.NewBadRequestApiError("Error al actualizar curso")
	}

	return course, nil
}

func DeleteCourse(course models.Course) e.ApiError {
	result := Db.Delete(&course).Where("Id_course = ?", course.Id_course)
	if result.Error != nil {
		log.Error("Error al eliminar el curso")
		log.Error(result.Error)
		return e.NewBadRequestApiError("Error al eliminar curso")
	}

	return nil
}
