package controller

import (
	"log"

	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type DefaultReceitaController interface {
	GetReceitaHandler(c *gin.Context)
}

func NewDefaultReceitaController(service service.ReceitaService) DefaultReceitaController {
	return &receitaController{
		service: service,
	}
}

func (ctrl *receitaController) GetReceitaHandler(c *gin.Context) {
	p1 := c.Param("p1")
	p2 := c.Param("p2")

	if p1 != "" && p2 != "" {
		ctrl.FindReceitaByAnoAndMes(c)
	} else if p1 != "" && p2 == "" {
		ctrl.FindReceitaById(c)
	} else {
		log.Println("Deu Ruim Familia!")
	}
}
