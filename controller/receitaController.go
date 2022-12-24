package controller

import (
	"log"
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/database"
	"github.com/LouisMatos/challenge-backend-2-go/model"
	"github.com/gin-gonic/gin"
)

func CadastraReceita(c *gin.Context) {
	var receita model.Receita
	log.Print("Iniciando cadastro de receita")

	if err := c.ShouldBindJSON(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	// date, _ := time.Parse("02/01/2006 15:04:05", receita.Data+" 00:00:00")

	// receita.Data = date.String()

	database.DB.Create(&receita)
	c.JSON(http.StatusOK, receita)
}

func BuscaTodasReceitas(c *gin.Context) {
	var receitas []model.Receita
	database.DB.Find(&receitas)
	c.JSON(200, receitas)
}
