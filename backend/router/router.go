package router

import (
	courses "backend/controllers/courses"
	users "backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)
	engine.POST("/users", users.RegisterUser)
	engine.PUT("/users", users.UpdateUser)

	engine.POST("/courses", courses.CreateCourse)
	engine.PUT("/courses", courses.UpdateCourse)
	engine.DELETE("/courses", courses.DeleteCourse)

}
