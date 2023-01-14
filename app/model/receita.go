package model

import (
	"time"

	"gopkg.in/validator.v2"
)

type Receita struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Descricao string    `gorm:"not null" json:"descricao"`
	Valor     float64   `gorm:"not null" json:"valor"`
	Data      time.Time `gorm:"not null" json:"data"`
}

type ReceitaDTO struct {
	Descricao string `json:"descricao" validate:"nonzero, nonnil"`
	Valor     string `json:"valor" validate:"nonzero, min=4, nonnil, regexp=^[\\d]+[.][\\d]{2}$"`
	Data      string `json:"data" validate:"nonzero, min=10, nonnil, regexp=^([0]?[1-9]|[1|2][0-9]|[3][0|1])[./-]([0]?[1-9]|[1][0-2])[./-]([0-9]{4}|[0-9]{2})$"`
}

func ValidaDadosReceita(receita *ReceitaDTO) error {
	if err := validator.Validate(receita); err != nil {
		return err
	}
	return nil
}
