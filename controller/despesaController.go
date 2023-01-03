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

func AtualizarDespesaPorID(c *gin.Context) {

	id := c.Params.ByName("id")

	despesa := service.BuscarDespesaId(id)

	if despesa.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Despesa não encontrada",
			"status":   404})
		return
	}

	var despesaDTO model.DespesaDTO

	log.Print("Iniciando atualização da despesa!")

	if err := c.ShouldBindJSON(&despesaDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	log.Print("Convertendo para json")

	if err := model.ValidaDadosDespesa(&despesaDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	log.Print("Campos validados com sucesso! ", despesaDTO)

	despesaAtualizada, jaCadastrado := service.AtualizarDespesa(&despesaDTO, id)

	if jaCadastrado {
		log.Println("Despesa cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Despesa já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, despesaAtualizada)
	}
}

func DeletarDespesaPorID(c *gin.Context) {

	id := c.Params.ByName("id")

	receita := service.BuscarDespesaId(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Despesa não encontrado",
			"status":   404})
		return
	}

	service.DeletarDespesaPorID(id)

	c.JSON(http.StatusNoContent, nil)
}

func BuscarDespesaId(c *gin.Context) {

	id := c.Params.ByName("p1")

	despesa := service.BuscarDespesaId(id)

	if despesa.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Despesa não encontrada!",
			"status":   404})
		return
	}

	c.JSON(http.StatusOK, despesa)
}

func BuscarTodasDespesas(c *gin.Context) {

	log.Println("Iniciando busca de todas as despesas!")

	descricao := c.Query("descricao")

	despesas := service.BuscarTodasDespesas(descricao, c)

	if len(despesas) == 0 {
		log.Println("Nenhuma despesa cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma despesa cadastrada!"})
	} else {
		c.JSON(200, despesas)
	}

}

func BuscarDespesaAnoMes(c *gin.Context) {

	mes := c.Params.ByName("p2")

	ano := c.Params.ByName("p1")

	despesas := service.BuscaTodasDespesasMesAno(mes, ano)

	if len(despesas) == 0 {
		log.Println("Nenhuma despesa cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma despesa cadastrada!"})
	} else {
		c.JSON(200, despesas)
	}
}
