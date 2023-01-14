package controller

import (
	"log"
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type DespesaController interface {
	Save(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindDespesaByAnoAndMes(c *gin.Context)
	FindDespesaById(c *gin.Context)
}

type despesaController struct {
	service service.DespesaService
}

func NewDespesaController(service service.DespesaService) DespesaController {
	return &despesaController{
		service: service,
	}
}

func (ctrl *despesaController) Save(c *gin.Context) {

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

	despesaSalva, jaCadastrado := ctrl.service.Save(despesa)

	if jaCadastrado {
		log.Println("Receita cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Receita já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, despesaSalva)
	}

}

func (ctrl *despesaController) GetAll(c *gin.Context) {

	log.Println("Iniciando busca de todas as despesas!")

	descricao := c.Query("descricao")

	despesas := ctrl.service.GetAll(descricao)

	if len(despesas) == 0 {
		log.Println("Nenhuma despesa cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma despesa cadastrada!"})
	} else {
		c.JSON(200, despesas)
	}
}

func (ctrl *despesaController) Update(c *gin.Context) {

	id := c.Params.ByName("id")

	despesa := ctrl.service.FindById(id)

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

	despesaAtualizada, jaCadastrado := ctrl.service.Update(despesaDTO, id)

	if jaCadastrado {
		log.Println("Despesa cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Despesa já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, despesaAtualizada)
	}

}

func (ctrl *despesaController) Delete(c *gin.Context) {

	id := c.Params.ByName("id")

	receita := ctrl.service.FindById(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Despesa não encontrado",
			"status":   404})
		return
	}

	ctrl.service.Delete(id)

	c.JSON(http.StatusNoContent, nil)

}

func (ctrl *despesaController) FindDespesaByAnoAndMes(c *gin.Context) {

	mes := c.Params.ByName("p2")

	ano := c.Params.ByName("p1")

	despesas := ctrl.service.FindDespesaByAnoAndMes(ano, mes)

	if len(despesas) == 0 {
		log.Println("Nenhuma despesa cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma despesa cadastrada!"})
	} else {
		c.JSON(200, despesas)
	}

}

func (ctrl *despesaController) FindDespesaById(c *gin.Context) {

	id := c.Params.ByName("p1")

	despesa := ctrl.service.FindById(id)

	if despesa.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Despesa não encontrada!",
			"status":   404})
		return
	}

	c.JSON(http.StatusOK, despesa)
}
