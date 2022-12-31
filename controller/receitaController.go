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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Receita já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, receitaSalva)
	}

}

func AtualizarReceitaPorID(c *gin.Context) {

	id := c.Params.ByName("id")

	receita := service.BuscarReceitaId(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Receita não encontrado",
			"status":   404})
		return
	}

	var receitaDTO model.ReceitaDTO

	log.Print("Iniciando atualização de receita")

	if err := c.ShouldBindJSON(&receitaDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	log.Print("Convertendo para json")

	if err := model.ValidaDadosReceita(&receitaDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	log.Print("Campos validados com sucesso! ", receitaDTO)

	receitaAtualizada, jaCadastrado := service.AtualizarReceita(&receitaDTO, id)

	if jaCadastrado {
		log.Println("Receita cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Receita já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, receitaAtualizada)
	}

}

func DeletarReceitaPorID(c *gin.Context) {

	id := c.Params.ByName("id")

	receita := service.BuscarReceitaId(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Receita não encontrado",
			"status":   404})
		return
	}

	service.DeletarReceitaPorID(id)

	c.JSON(http.StatusNoContent, nil)
}

func BuscarReceitaId(c *gin.Context) {

	id := c.Params.ByName("p1")
	log.Println(id)

	receita := service.BuscarReceitaId(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Receita não encontrado",
			"status":   404})
		return
	}

	c.JSON(200, receita)

}

func BuscaTodasReceitas(c *gin.Context) {

	log.Println("Iniciando busca de todas as receitas!")

	descricao := c.Query("descricao")

	receitas := service.BuscaTodasReceitas(descricao, c)

	if len(receitas) == 0 {
		log.Println("Nenhuma receita cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma receita cadastrada!"})
	} else {
		c.JSON(200, receitas)
	}
}

func BuscarReceitaAnoMes(c *gin.Context) {

	mes := c.Params.ByName("p2")

	ano := c.Params.ByName("p1")

	receitas := service.BuscaTodasReceitasMesAno(mes, ano)

	if len(receitas) == 0 {
		log.Println("Nenhuma receita cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma receita cadastrada!"})
	} else {
		c.JSON(200, receitas)
	}
}
