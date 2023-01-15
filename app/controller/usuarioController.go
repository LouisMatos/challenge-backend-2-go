package controller

import (
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type UsuarioController interface {
	Save(c *gin.Context)
}

type usuarioController struct {
	service service.UsuarioService
}

func NewUsuarioController(service service.UsuarioService) UsuarioController {
	return &usuarioController{
		service: service,
	}
}

func (ctrl *usuarioController) Save(c *gin.Context) {

	var usuario model.Usuario

	err := c.ShouldBindJSON(&usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	_, err = ctrl.service.Save(usuario)

	if err != nil {
		_, err = ctrl.service.Save(usuario)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
