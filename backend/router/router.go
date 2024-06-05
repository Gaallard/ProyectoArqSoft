package router

import (
	courses "backend/controllers/courses"
	inscriptions "backend/controllers/inscriptions"
	users "backend/controllers/users"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)
	engine.POST("/users/register", users.RegisterUser)

	authorized := engine.Group("")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/inscriptions", inscriptions.UserInscription)
		authorized.GET("/courses/:idUser", courses.GetCoursesByUser)
		admin := authorized.Group("")
		admin.Use(middleware.AuthMiddlewareAdmin())
		{
			admin.POST("/courses/create", courses.CreateCourse)
			admin.PUT("/courses/update", courses.UpdateCourse)
			admin.DELETE("/courses/delete", courses.DeleteCourse)
		}
	}
}
