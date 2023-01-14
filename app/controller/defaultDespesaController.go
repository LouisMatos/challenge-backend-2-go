package controller

import (
	"log"

	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type DefaultDespesaController interface {
	GetDespesaHandler(c *gin.Context)
}

func NewDefaultDespesaController(service service.DespesaService) DefaultDespesaController {
	return &despesaController{
		service: service,
	}
}

func (ctrl *despesaController) GetDespesaHandler(c *gin.Context) {
	p1 := c.Param("p1")
	p2 := c.Param("p2")

	if p1 != "" && p2 != "" {
		ctrl.FindDespesaByAnoAndMes(c)
	} else if p1 != "" && p2 == "" {
		ctrl.FindDespesaById(c)
	} else {
		log.Println("Deu Ruim Familia!")
	}
}
