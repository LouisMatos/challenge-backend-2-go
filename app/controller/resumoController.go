package controller

import (
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

func RealizarResumoAnoMes(c *gin.Context) {

	mes := c.Params.ByName("mes")

	ano := c.Params.ByName("ano")

	buscaResumo := service.RealizaResumoAnoMes(mes, ano)

	c.JSON(200, buscaResumo)

}
