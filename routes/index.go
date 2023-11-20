package routes

import "github.com/gin-gonic/gin"

func AddRoutes(g *gin.RouterGroup) {
	authRoutes(g)
}
