package controller

import (
	"log"
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
)

type ReceitaController interface {
	Save(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindReceitaByAnoAndMes(c *gin.Context)
	FindReceitaById(c *gin.Context)
}

type receitaController struct {
	service service.ReceitaService
}

func NewReceitaController(service service.ReceitaService) ReceitaController {
	return &receitaController{
		service: service,
	}
}

func (ctrl *receitaController) FindReceitaByAnoAndMes(c *gin.Context) {

	mes := c.Params.ByName("p2")

	ano := c.Params.ByName("p1")

	receitas := ctrl.service.FindReceitaByAnoAndMes(ano, mes)

	if len(receitas) == 0 {
		log.Println("Nenhuma receita cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma receita cadastrada!"})
	} else {
		c.JSON(200, receitas)
	}
}

func (ctrl *receitaController) FindReceitaById(c *gin.Context) {

	id := c.Params.ByName("p1")
	log.Println(id)

	receita := ctrl.service.FindById(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Receita não encontrado",
			"status":   404})
		return
	}

	c.JSON(200, receita)
}

func (ctrl *receitaController) Save(c *gin.Context) {

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

	receitaSalva, jaCadastrado := ctrl.service.Save(receita)

	if jaCadastrado {
		log.Println("Receita cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Receita já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, receitaSalva)
	}
}

func (ctrl *receitaController) Update(c *gin.Context) {

	id := c.Params.ByName("id")

	receita := ctrl.service.FindById(id)

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

	receitaAtualizada, jaCadastrado := ctrl.service.Update(receitaDTO, id)

	if jaCadastrado {
		log.Println("Receita cadastrada anteriormente!")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Receita já cadastrada nesse mês!", "status": 422})
	} else {
		c.JSON(http.StatusOK, receitaAtualizada)
	}

}

func (ctrl *receitaController) Delete(c *gin.Context) {

	id := c.Params.ByName("id")

	receita := ctrl.service.FindById(id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Receita não encontrado",
			"status":   404})
		return
	}

	ctrl.service.Delete(id)

	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *receitaController) GetAll(c *gin.Context) {

	log.Println("Iniciando busca de todas as receitas!")

	descricao := c.Query("descricao")

	receitas := ctrl.service.GetAll(descricao)

	if len(receitas) == 0 {
		log.Println("Nenhuma receita cadastrada!")
		c.JSON(404, gin.H{"status": 404, "mensagem": "Nenhuma receita cadastrada!"})
	} else {
		c.JSON(200, receitas)
	}

}
