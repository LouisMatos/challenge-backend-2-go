package controller

import (
	"log"
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/model"
	"github.com/LouisMatos/challenge-backend-2-go/service"
	"github.com/gin-gonic/gin"
)

func CadastraReceita(c *gin.Context) {

	var receita model.ReceitaDTO
	log.Print("Iniciando cadastro de receita")

	if err := c.ShouldBindJSON(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	log.Print("Convertendo para json")

	if err := model.ValidaDadosReceita(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	log.Print("Campos validados com sucesso! ", receita)

	receitaSalva, jaCadastrado := service.SalvarNovaReceita(&receita, c)

	if jaCadastrado {
		log.Println("Receita cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"erro": "Receita já cadastrada nesse mês!"})
	} else {
		c.JSON(http.StatusOK, receitaSalva)
	}

}

func BuscaTodasReceitas(c *gin.Context) {
	log.Println("Iniciando busca de todas as receitas!")
	receitas := service.BuscaTodasReceitas(c)
	c.JSON(200, receitas)
}
