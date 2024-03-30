package main

import (
	"os"

	"github.com/NetKBs/uneg-sase-go-backend/api/routes"
	"github.com/NetKBs/uneg-sase-go-backend/initializers"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// @title           UNEG Services API
// @version         1.0
// @description     This is an api to manage students, teachers and academic processes in the university
// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	PORT := ":" + os.Getenv("PORT_SERVER")

	router := gin.Default()
	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.SetupUsersRoutes(router)
	router.Run(PORT)
}
