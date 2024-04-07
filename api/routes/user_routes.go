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

		usersGroup.POST("/create", user.CreateUser)
		usersGroup.GET("/getall", user.GetAllUser)

		rolesGroup := usersGroup.Group("/roles")
		{
			rolesGroup.POST("/create", user.CreateRole)
			rolesGroup.PUT("/update", user.UpdateRole)
			rolesGroup.DELETE("/delete/:id", user.DeleteRole)
			rolesGroup.GET("/getall", user.GetAllRole)
			rolesGroup.GET("/get/:id", user.GetRole)
		}

	}
}
