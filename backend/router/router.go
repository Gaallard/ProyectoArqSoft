package router

import (
	courses "backend/controllers/courses"
	users "backend/controllers/users"
	middleware "backend/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)

	authorized := engine.Group("")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/users", users.RegisterUser)
		authorized.PUT("/users", users.UpdateUser)
		authorized.POST("/courses", courses.CreateCourse)
		authorized.PUT("/courses", courses.UpdateCourse)
		authorized.DELETE("/courses", courses.DeleteCourse)
	}

}
