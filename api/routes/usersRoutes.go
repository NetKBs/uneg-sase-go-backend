package routes

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/controllers"
	"github.com/NetKBs/uneg-sase-go-backend/api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", controllers.Singup)
		userGroup.POST("/login", controllers.Login)
		userGroup.GET("/validated", middlewares.RequireAuth, controllers.Validated)
	}
}
