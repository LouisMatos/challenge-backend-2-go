package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/app/config"
	"github.com/LouisMatos/challenge-backend-2-go/app/controller"
	"github.com/LouisMatos/challenge-backend-2-go/app/database"
	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func SetupDatabase() {
	config.LoadAppConfig()
	database.ConexaoComBancoDados(config.AppConfig.ConnectionString)
	database.Migrate()
}

func CriarReceitaMock() {
	receita := model.Receita{Descricao: "Descrição Teste", Valor: 12.21, Data: time.Now()}
	database.DB.Create(&receita)
	ID = int(receita.ID)
	log.Println("Salva Receita Mock", receita)
}

func CriarReceitaDtoMock() model.ReceitaDTO {
	return model.ReceitaDTO{Descricao: "Descrição Teste", Valor: "12.21", Data: "12/10/2022"}
}

func CriarReceitaDtoMockInvalido() model.ReceitaDTO {
	return model.ReceitaDTO{Descricao: "", Valor: "", Data: ""}
}

func DeletaReceitaMock() {
	var receita model.Receita
	database.DB.Delete(&receita, ID)
	log.Println("Deleta Receita Mock ID", ID)
}

func TestCriaNovaReceita(t *testing.T) {
	SetupDatabase()
	CriarReceitaMock()
	defer DeletaReceitaMock()

	r := SetupDasRotasDeTeste()
	r.POST("/receitas", controller.CadastraReceita)

	receita := CriarReceitaDtoMock()
	valorJson, _ := json.Marshal(receita)
	req, _ := http.NewRequest("POST", "/receitas", bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var receitaMock model.Receita
	json.Unmarshal(resposta.Body.Bytes(), &receitaMock)
	assert.Equal(t, "Descrição Teste", receitaMock.Descricao)
	assert.Equal(t, float64(12.21), receitaMock.Valor)
	assert.Equal(t, "2022-10-12 00:00:00 +0000 UTC", receitaMock.Data.String())

}

func TestCriaReceitaCamposInvalidos(t *testing.T) {

	r := SetupDasRotasDeTeste()
	r.POST("/receitas", controller.CadastraReceita)

	receita := CriarReceitaDtoMockInvalido()
	valorJson, _ := json.Marshal(receita)
	req, _ := http.NewRequest("POST", "/receitas", bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, "{\"erro\":\"Descricao: zero value, Valor: zero value, less than min, regular expression mismatch, Data: zero value, less than min, regular expression mismatch\"}", resposta.Body.String())

}

func TestCriaReceitaJaCadastrada(t *testing.T) {
	SetupDatabase()
	CriarReceitaMock()
	defer DeletaReceitaMock()

	r := SetupDasRotasDeTeste()
	r.POST("/receitas", controller.CadastraReceita)

	receita := CriarReceitaDtoMock()
	valorJson, _ := json.Marshal(receita)
	req, _ := http.NewRequest("POST", "/receitas", bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, 200, resposta.Result().StatusCode)

	req2, _ := http.NewRequest("POST", "/receitas", bytes.NewBuffer(valorJson))
	resposta2 := httptest.NewRecorder()
	r.ServeHTTP(resposta2, req2)

	assert.Equal(t, "{\"message\":\"Receita já cadastrada nesse mês!\",\"status\":422}", resposta2.Body.String())

}
