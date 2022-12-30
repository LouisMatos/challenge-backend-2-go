package controller

import (
	"log"
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/model"
	"github.com/LouisMatos/challenge-backend-2-go/service"
	"github.com/gin-gonic/gin"
)

func CadastrarDespesa(c *gin.Context) {

	var despesa model.DespesaDTO
	log.Print("Iniciando cadastro de receita")

	if err := c.ShouldBindJSON(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	log.Print("Convertendo para json")

	if err := model.ValidaDadosDespesa(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	log.Print("Campos validados com sucesso! ", despesa)

	despesaSalva, jaCadastrado := service.SalvarNovaDespesa(&despesa, c)

	if jaCadastrado {
		log.Println("Receita cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Receita já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, despesaSalva)
	}

}

func BuscarTodasDespesas(c *gin.Context) {

	log.Println("Iniciando busca de todas as despesas!")

	despesas := service.BuscarTodasDespesas(c)

	if len(despesas) == 0 {
		log.Println("Nenhuma despesa cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma despesa cadastrada!"})
	} else {
		c.JSON(200, despesas)
	}

}
