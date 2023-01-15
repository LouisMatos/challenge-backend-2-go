package model

import (
	"time"
)

type Usuario struct {
	Id        uint64    `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"not null" json:"-" time_format:"2006-01-02" time_utc:"1"`
	IsActive  int       `gorm:"not null" json:"is_active"`
}
