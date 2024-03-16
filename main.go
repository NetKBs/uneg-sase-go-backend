package main

import (
	"os"

	"github.com/NetKBs/uneg-sase-go-backend/api/routes"
	"github.com/NetKBs/uneg-sase-go-backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	PORT := ":" + os.Getenv("PORT_SERVER")

	router := gin.Default()
	routes.SetupUsersRoutes(router)
	router.Run(PORT)
}
