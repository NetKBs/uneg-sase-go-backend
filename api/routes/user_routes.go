package routes

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/controllers/user"
	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(router *gin.Engine) {

	userGroup := router.Group("/users")
	{

		userGroup.GET("/")
		//userGroup.POST("/signup", controllers.Singup)
		//userGroup.POST("/login", controllers.Login)
		//userGroup.GET("/validated", middlewares.RequireAuth, controllers.Validated)

		//userGroup.POST("/roles/create", controllers.RoleCreate)
		userGroup.POST("roles/create", user.Create)
	}
}
