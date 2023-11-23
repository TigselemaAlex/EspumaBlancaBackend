package controllers

import (
	"EspumaBlancaBackend/data/request"
	"EspumaBlancaBackend/data/response"
	"EspumaBlancaBackend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleController struct {
	roleService service.RoleService
}

func NewRoleController(service service.RoleService) *RoleController {
	return &RoleController{roleService: service}
}

func (controller *RoleController) Create(ctx *gin.Context) {
	roleRequest := request.RoleRequest{}
	err := ctx.ShouldBindJSON(&roleRequest)
	if err != nil {
		panic(err)
	}

	controller.roleService.Create(roleRequest)

	apiResponse := response.ApiResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   nil,
	}

	ctx.JSON(apiResponse.Code, apiResponse)
}

func (controller *RoleController) FindById(ctx *gin.Context) {
	roleId := ctx.Param("id")
	id, err := strconv.Atoi(roleId)

	if err != nil {
		panic(err)
	}

	serviceResponse := controller.roleService.FindById(uint(id))

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   serviceResponse,
	}

	ctx.JSON(apiResponse.Code, apiResponse)
}

func (controller *RoleController) FindAll(ctx *gin.Context) {
	serviceResponse := controller.roleService.FindAll()

	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   serviceResponse,
	}

	ctx.JSON(apiResponse.Code, apiResponse)
}
