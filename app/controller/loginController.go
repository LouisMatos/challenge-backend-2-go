package controller

import (
	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	usuarioService service.UsuarioService
	jWtService     service.JWTService
}

func NewLoginController(usuarioService service.UsuarioService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		usuarioService: usuarioService,
		jWtService:     jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var usuario model.Usuario
	err := ctx.ShouldBind(&usuario)
	if err != nil {
		return ""
	}

	isAuthenticated := controller.usuarioService.Login(usuario.Email, usuario.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(usuario.Email, true)
	}
	return ""
}
