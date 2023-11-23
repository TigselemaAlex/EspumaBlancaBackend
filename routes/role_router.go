package routes

import (
	"EspumaBlancaBackend/config"
	"EspumaBlancaBackend/controllers"
	"EspumaBlancaBackend/repository"
	"EspumaBlancaBackend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func roleRoutes(g *gin.RouterGroup) {
	db := config.DatabaseConnection()
	validate := validator.New()
	roleRepository := repository.NewRoleRepositoryImpl(db)
	roleService := service.NewRoleServiceImpl(roleRepository, validate)
	roleController := controllers.NewRoleController(roleService)
	r := g.Group("/public/roles")
	{
		r.POST("/", roleController.Create)
		r.GET("/", roleController.FindAll)
		r.GET("/:id", roleController.FindById)
	}

}
