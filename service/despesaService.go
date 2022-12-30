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

	var despesa model.Despesa

	database.DB.Where("descricao ILIKE ? AND TO_CHAR(data, 'yyyy-mm') LIKE ?", Descricao, Data.Format("2006-01")).Find(&despesa)

	if despesa.ID == 0 {
		log.Println("Despesa ainda não foi cadastrada!")
		return false
	} else {
		log.Println("Despesa já foi cadastrada!")
		return true
	}

}

func BuscarDespesaId(id string) model.Despesa {
	var despesa model.Despesa
	database.DB.First(&despesa, id)
	return despesa
}

func DeletarDespesaPorID(id string) {
	var despesa model.Despesa
	database.DB.Delete(&despesa, id)
}

func AtualizarDespesa(despesaDTO *model.DespesaDTO, id string) (model.Despesa, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", despesaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(despesaDTO.Valor, 32)

	u, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		panic(err)
	}

	despesa := model.Despesa{
		ID:        uint(u),
		Descricao: despesaDTO.Descricao,
		Data:      date,
		Valor:     float32(value),
	}

	isSaved := validarDespesaJaCadastrada(despesa.Descricao, despesa.Data)

	if !isSaved {

		log.Println("Convertendo dto para objeto a ser atualizado no banco de dados!")

		database.DB.Save(&despesa)

		log.Println("Despesa atualizada no banco de dados!")

		return despesa, false
	} else {
		return despesa, true
	}
}

func BuscarTodasDespesas(c *gin.Context) []model.Despesa {

	var despesas []model.Despesa

	database.DB.Find(&despesas)

	return despesas

}
