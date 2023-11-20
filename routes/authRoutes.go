package routes

import (
	"EspumaBlancaBackend/controllers"
	"github.com/gin-gonic/gin"
)

func authRoutes(g *gin.RouterGroup) {
	rAuth := g.Group("/auth")
	{
		rAuth.POST("/signup", controllers.Signup)
		rAuth.POST("/login", controllers.Login)
	}
}
