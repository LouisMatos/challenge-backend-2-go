package model

import (
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/enum"
	"gopkg.in/validator.v2"
)

type Despesa struct {
	ID        uint           `gorm:"primary_key"`
	Descricao string         `gorm:"not null"`
	Valor     float32        `gorm:"not null"`
	Data      time.Time      `gorm:"not null"`
	Categoria enum.Categoria `gorm:"not null"`
}

type DespesaDTO struct {
	Descricao string `json:"descricao" validate:"nonzero, nonnil"`
	Valor     string `json:"valor" validate:"nonzero, min=4, nonnil, regexp=^[\\d]+[.][\\d]{2}$"`
	Data      string `json:"data" validate:"nonzero, min=10, nonnil, regexp=^([0]?[1-9]|[1|2][0-9]|[3][0|1])[./-]([0]?[1-9]|[1][0-2])[./-]([0-9]{4}|[0-9]{2})$"`
	Categoria string `json:"categoria"`
}

func ValidaDadosDespesa(despesa *DespesaDTO) error {
	if err := validator.Validate(despesa); err != nil {
		return err
	}
	return nil
}

// var d enum.Categoria = 2
// 	fmt.Println(d)             // Print : West
// 	fmt.Println(d.String())    // Print : West
// 	fmt.Println(d.EnumIndex()) // Print : 4
