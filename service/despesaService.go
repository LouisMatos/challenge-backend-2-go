package service

import (
	"log"
	"strconv"
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/database"
	"github.com/LouisMatos/challenge-backend-2-go/model"
	"github.com/gin-gonic/gin"
)

func SalvarNovaDespesa(despesaDTO *model.DespesaDTO, c *gin.Context) (model.Despesa, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", despesaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(despesaDTO.Valor, 32)

	despesa := model.Despesa{
		Descricao: despesaDTO.Descricao,
		Data:      date,
		Valor:     float32(value),
	}

	isSaved := validarDespesaJaCadastrada(despesa.Descricao, despesa.Data)

	if !isSaved {

		log.Println("Convertendo dto para objeto a ser salvo no banco de dados!")

		database.DB.Create(&despesa)

		log.Println("Despesa salva no banco de dados!")

		return despesa, false

	} else {
		return despesa, true
	}

}

func validarDespesaJaCadastrada(Descricao string, Data time.Time) bool {

	var receita model.Receita

	database.DB.Where("descricao ILIKE ? AND TO_CHAR(data, 'yyyy-mm') LIKE ?", Descricao, Data.Format("2006-01")).Find(&receita)

	if receita.ID == 0 {
		log.Println("Despesa ainda não foi cadastrada!")
		return false
	} else {
		log.Println("Despesa já foi cadastrada!")
		return true
	}

}

func BuscarTodasDespesas(c *gin.Context) []model.Despesa {

	var despesas []model.Despesa

	database.DB.Find(&despesas)

	return despesas

}
