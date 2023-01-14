package controller

import (
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type ResumoController interface {
	GetMonthSummary(c *gin.Context)
}

type resumoController struct {
	service service.ResumoService
}

func NewResumoController(service service.ResumoService) ResumoController {
	return &resumoController{
		service: service,
	}
}

func (ctrl *resumoController) GetMonthSummary(c *gin.Context) {

	mes := c.Params.ByName("mes")

	ano := c.Params.ByName("ano")

	buscaResumo := ctrl.service.GetMonthSummary(ano, mes)

	c.JSON(200, buscaResumo)
}
