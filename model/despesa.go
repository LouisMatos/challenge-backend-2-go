package model

import (
	"time"

	"gopkg.in/validator.v2"
)

type Despesa struct {
	ID        uint      `gorm:"primary_key"`
	Descricao string    `gorm:"not null"`
	Valor     float32   `gorm:"not null"`
	Data      time.Time `gorm:"not null"`
}

type DespesaDTO struct {
	Descricao string `json:"descricao" validate:"nonzero, nonnil"`
	Valor     string `json:"valor" validate:"nonzero, min=4, nonnil, regexp=^[\\d]+[.][\\d]{2}$"`
	Data      string `json:"data" validate:"nonzero, min=10, nonnil, regexp=^([0]?[1-9]|[1|2][0-9]|[3][0|1])[./-]([0]?[1-9]|[1][0-2])[./-]([0-9]{4}|[0-9]{2})$"`
}

func ValidaDadosDespesa(despesa *DespesaDTO) error {
	if err := validator.Validate(despesa); err != nil {
		return err
	}
	return nil
}
