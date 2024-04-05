package routes

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/controllers/user"
	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(router *gin.Engine) {

	usersGroup := router.Group("/users")
	{

		usersGroup.GET("/")
		//userGroup.POST("/signup", controllers.Singup)
		//userGroup.POST("/login", controllers.Login)
		//userGroup.GET("/validated", middlewares.RequireAuth, controllers.Validated)

		//userGroup.POST("/roles/create", controllers.RoleCreate)

		rolesGroup := usersGroup.Group("/roles")
		{
			rolesGroup.POST("/create", user.Create)
			rolesGroup.PUT("/update", user.Update)
			rolesGroup.DELETE("/delete/:id", user.Delete)
			rolesGroup.GET("/getall", user.GetAll)
			rolesGroup.GET("/get/:id", user.Get)
		}

	}
}
