package service

import (
	"math"

	"github.com/LouisMatos/challenge-backend-2-go/database"
	"github.com/LouisMatos/challenge-backend-2-go/model"
)

func RealizaResumoAnoMes(mes string, ano string) model.Resumo {

	var resumo model.Resumo

	var despesas []model.Despesa
	var receitas []model.Receita

	if len(mes) == 1 {
		mes = "0" + mes
	}

	database.DB.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&despesas)

	database.DB.Where("(TO_CHAR(data, 'YYYY-MM')) = ?", ""+ano+"-"+mes).Find(&receitas)

	for i := 0; i < len(despesas); i++ {
		resumo.ValorTotalDespesa = resumo.ValorTotalDespesa + float64(despesas[i].Valor)
	}

	for i := 0; i < len(receitas); i++ {
		resumo.ValorTotalReceita = resumo.ValorTotalDespesa + float64(receitas[i].Valor)
	}

	resumo.ValorTotalDespesa = toFixed(resumo.ValorTotalDespesa, 2)
	resumo.ValorTotalReceita = toFixed(resumo.ValorTotalReceita, 2)
	resumo.SaldoFinal = toFixed(resumo.ValorTotalReceita-resumo.ValorTotalDespesa, 2)

	return resumo

}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
