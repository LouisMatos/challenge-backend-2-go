package service

import (
	"log"
	"strconv"
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/database"
	"github.com/LouisMatos/challenge-backend-2-go/model"
	"github.com/gin-gonic/gin"
)

func SalvarNovaReceita(receitaDTO *model.ReceitaDTO, c *gin.Context) (model.Receita, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", receitaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(receitaDTO.Valor, 32)

	receita := model.Receita{
		Descricao: receitaDTO.Descricao,
		Data:      date,
		Valor:     float32(value),
	}

	isSaved := validarReceitaJaCadastrada(receita.Descricao, receita.Data)

	if !isSaved {

		log.Println("Convertendo dto para objeto a ser salvo no banco de dados!")

		database.DB.Create(&receita)

		log.Println("Receita salva no banco de dados!")

		return receita, false

	} else {
		return receita, true
	}

}

func AtualizarReceita(receitaDTO *model.ReceitaDTO, id string) (model.Receita, bool) {

	date, _ := time.Parse("02/01/2006 15:04:05", receitaDTO.Data+" 00:00:00")

	value, _ := strconv.ParseFloat(receitaDTO.Valor, 32)

	u, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		panic(err)
	}

	receita := model.Receita{
		ID:        uint(u),
		Descricao: receitaDTO.Descricao,
		Data:      date,
		Valor:     float32(value),
	}

	isSaved := validarReceitaJaCadastrada(receita.Descricao, receita.Data)

	if !isSaved {

		log.Println("Convertendo dto para objeto a ser atualizado no banco de dados!")

		database.DB.Save(&receita)

		log.Println("Receita atualizada no banco de dados!")

		return receita, false
	} else {
		return receita, true
	}

}

func validarReceitaJaCadastrada(Descricao string, Data time.Time) bool {

	var receita model.Receita

	database.DB.Where("descricao ILIKE ? AND TO_CHAR(data, 'yyyy-mm') LIKE ?", Descricao, Data.Format("2006-01")).Find(&receita)

	if receita.ID == 0 {
		log.Println("Receita ainda não foi cadastrada!")
		return false
	} else {
		log.Println("Receita já foi cadastrada!")
		return true
	}

}

func DeletarReceitaPorID(id string) {
	var receita model.Receita
	database.DB.Delete(&receita, id)
}

func BuscarReceitaId(id string) model.Receita {
	var receita model.Receita
	database.DB.First(&receita, id)
	return receita
}

func BuscaTodasReceitas(descricao string, c *gin.Context) []model.Receita {

	var receitas []model.Receita

	database.DB.Where("descricao ILIKE ?", "%"+descricao+"%").Find(&receitas)

	return receitas

}

func BuscaTodasReceitasMesAno(mes string, ano string) []model.Receita {

	var receitas []model.Receita

	if len(mes) == 1 {
		mes = "0" + mes
	}

	database.DB.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&receitas)

	return receitas

}
