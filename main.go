package main

import (
	"EspumaBlancaBackend/controllers"
	"EspumaBlancaBackend/initializers"
	"EspumaBlancaBackend/middleware"
	"EspumaBlancaBackend/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	g := r.Group(os.Getenv("CONTEXT_PATH"))
	routes.AddRoutes(g)
	g.GET("/validate", middleware.RequireAuth, controllers.ValidateToken)
	r.Run()
}
