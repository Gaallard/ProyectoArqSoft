package router

import (
	users "backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)
}
