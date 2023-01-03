package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetReceitaHandler(c *gin.Context) {
	p1 := c.Param("p1")
	p2 := c.Param("p2")

	if p1 != "" && p2 != "" {
		BuscarReceitaAnoMes(c)
	} else if p1 != "" && p2 == "" {
		BuscarReceitaId(c)
	} else {
		log.Println("Deu Ruim Familia!")
	}
}

func GetDespesaHandler(c *gin.Context) {
	p1 := c.Param("p1")
	p2 := c.Param("p2")

	if p1 != "" && p2 != "" {
		BuscarDespesaAnoMes(c)
	} else if p1 != "" && p2 == "" {
		BuscarDespesaId(c)
	} else {
		log.Println("Deu Ruim Familia!")
	}
}
